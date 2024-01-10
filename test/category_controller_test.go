package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang_restfulapi/app"
	"golang_restfulapi/controller"
	"golang_restfulapi/helper"
	"golang_restfulapi/middleware"
	"golang_restfulapi/model/entity"
	"golang_restfulapi/repository"
	"golang_restfulapi/services"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupDB() *sql.DB {
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/golang_restfulapi_test")
	helper.IfErrorPanic(err)

	db.SetMaxIdleConns(10)  // minimum number of connection
	db.SetMaxOpenConns(100) // maximum number of connection
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	return middleware.NewAuthMiddleware(router)
}

func emptyTable(db *sql.DB) {
	db.Exec("TRUNCATE categories")
}

// ========================================================= Create =========================================================

func TestCreateCategorySuccess(t *testing.T) {
	db := setupDB()
	emptyTable(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"Flora"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Flora", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupDB()
	emptyTable(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	// body, _ := io.ReadAll(response.Body)

	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
	// assert.Equal(t, "Flora", responseBody["data"].(map[string]interface{})["name"])
}

// ========================================================= Update =========================================================

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupDB()
	emptyTable(db)

	// create data first
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository().Create(context.Background(), tx, entity.Category{
		Name: "Gabriella",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"Cornelia"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(categoryRepository.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, categoryRepository.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Cornelia", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupDB()
	emptyTable(db)

	// create data first
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository().Create(context.Background(), tx, entity.Category{
		Name: "Gabriella",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(categoryRepository.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
	assert.Equal(t, categoryRepository.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	// assert.Equal(t, "Cornelia", responseBody["data"].(map[string]interface{})["name"])

}

// ========================================================= Delete By ID =========================================================

func TestDeleteByIdSuccess(t *testing.T) {
	db := setupDB()
	emptyTable(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository().Create(context.Background(), tx, entity.Category{
		Name: "Fiony",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(categoryRepository.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	// assert.Equal(t, categoryRepository.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	// assert.Equal(t, "Fiony", responseBody["data"].(map[string]interface{})["name"])
}

func TestDeleteByIdFailed(t *testing.T) {
	db := setupDB()
	emptyTable(db)

	// to simulate code 404 not found, doesn't need create data
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
}

// ========================================================= Find By ID =========================================================

func TestFindByIdSuccess(t *testing.T) {
	db := setupDB()
	emptyTable(db)

	// create data first
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository().Create(context.Background(), tx, entity.Category{
		Name: "Gabriella",
	})
	tx.Commit()

	router := setupRouter(db)

	// requestBody := strings.NewReader(`{"name":"Cornelia"}`)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(categoryRepository.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, categoryRepository.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, categoryRepository.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestFindByIdFailed(t *testing.T) {
	db := setupDB()
	emptyTable(db)

	// // create data first
	// tx, _ := db.Begin()
	// categoryRepository := repository.NewCategoryRepository().Create(context.Background(), tx, entity.Category{
	// 	Name: "",
	// })
	// tx.Commit()

	router := setupRouter(db)

	// requestBody := strings.NewReader(`{"name":"Cornelia"}`)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
	// assert.Equal(t, categoryRepository.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	// assert.Equal(t, categoryRepository.Name, responseBody["data"].(map[string]interface{})["name"])
}

// ========================================================= Find All =========================================================

func TestFindAllSuccess(t *testing.T) {
	db := setupDB()
	emptyTable(db)

	// create data first
	tx, _ := db.Begin()
	categoryRepository1 := repository.NewCategoryRepository().Create(context.Background(), tx, entity.Category{
		Name: "Gabriella",
	})
	categoryRepository2 := repository.NewCategoryRepository().Create(context.Background(), tx, entity.Category{
		Name: "Cornelia",
	})
	tx.Commit()

	router := setupRouter(db)

	// requestBody := strings.NewReader(`{"name":"Cornelia"}`)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(responseBody)

	// var categoriesDatas = responseBody["data"].([]map[string]interface{})

	var categoriesDatas = responseBody["data"].([]interface{})

	categoryData1 := categoriesDatas[0].(map[string]interface{})
	categoryData2 := categoriesDatas[1].(map[string]interface{})

	assert.Equal(t, categoryRepository1.Id, int(categoryData1["id"].(float64)))
	assert.Equal(t, categoryRepository1.Name, categoryData1["name"])

	assert.Equal(t, categoryRepository2.Id, int(categoryData2["id"].(float64)))
	assert.Equal(t, categoryRepository2.Name, categoryData2["name"])
}

func TestUnauthorized(t *testing.T) {
	db := setupDB()
	emptyTable(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("API-Key", "hahaha") // wrong key

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"])
}
