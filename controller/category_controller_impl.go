package controller

import (
	"golang_restfulapi/helper"
	"golang_restfulapi/model/web"
	"golang_restfulapi/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// get request body use json

	// create data to decode
	categoryCreateRequest := web.CategoryCreateRequest{} // create empty interface
	helper.ReadRequestBody(request, &categoryCreateRequest)

	// send response
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	// write data response into writer
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId") // return string, need convertion
	id, err := strconv.Atoi(categoryId)
	helper.IfErrorPanic(err)
	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) DeleteById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.IfErrorPanic(err)

	controller.CategoryService.DeleteById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK", // no return data
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) DeleteAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	controller.CategoryService.DeleteAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.IfErrorPanic(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindALl(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryResponses := controller.CategoryService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteResponseBody(writer, webResponse)
}
