package main

import (
	"golang_restfulapi/helper"
	"golang_restfulapi/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	server := InitializedServer()

	err := server.ListenAndServe()
	helper.IfErrorPanic(err)
}
