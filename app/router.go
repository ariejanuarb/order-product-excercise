package app

import (
	"github.com/julienschmidt/httprouter"
	"golang-rest-api/controller"
	"golang-rest-api/exception"
)

func NewCategoryRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func NewProductRouter(controller controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/products", controller.FindAll)
	router.GET("/api/products/:productId", controller.FindById)
	router.POST("/api/products", controller.Create)
	router.PUT("/api/products/:productId", controller.Update)
	router.DELETE("/api/products/:productId", controller.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
