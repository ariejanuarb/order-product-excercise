package main

import (
	"github.com/go-playground/validator/v10"
	"golang-rest-api/app"
	"golang-rest-api/controller"
	"golang-rest-api/helper"
	"golang-rest-api/middleware"
	"golang-rest-api/repository"
	"golang-rest-api/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()

	// memanggin contructor service category
	// dependency : categoryRepository, db, validate
	categoryService := service.NewCategoryService(categoryRepository, db, validate)

	// membuat object categoryController, dengan memanggil constructor new Category
	// depencency : category service
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
