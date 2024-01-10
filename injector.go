//go:build wireinject
// +build wireinject

package main

import (
	"golang_restfulapi/app"
	"golang_restfulapi/controller"
	"golang_restfulapi/middleware"
	"golang_restfulapi/repository"
	"golang_restfulapi/services"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		repository.NewCategoryRepository,
		services.NewCategoryService,
		controller.NewCategoryController,
		app.NewRouter, // return *httprouter.Router, need to bind
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
