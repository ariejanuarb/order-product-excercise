package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
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

	// memanggin contructor service categories
	// dependency : categoryRepository, db, validate
	categoryService := service.NewCategoryService(categoryRepository, db, validate)

	// membuat object categoryController, dengan memanggil constructor new Category
	// depencency : categories service
	categoryController := controller.NewCategoryController(categoryService)

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router := httprouter.New()

	app.NewCategoryRouter(router, categoryController)
	app.NewProductRouter(router, productController)

	cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://foo.com:8080"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	server := http.Server{
		Addr:    "localhost:3001",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
