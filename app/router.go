package app

import (
	"golang_restfulapi/controller"
	"golang_restfulapi/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindALl)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.DeleteById)
	router.DELETE("/api/categories", categoryController.DeleteAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
