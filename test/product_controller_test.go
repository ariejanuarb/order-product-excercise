package test

import (
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"golang-rest-api/controller"
	"golang-rest-api/exception"
	"golang-rest-api/middleware"
	"golang-rest-api/repository"
	"golang-rest-api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SetupProductRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	controller := controller.NewProductController(productService)

	router := httprouter.New()
	router.GET("/api/products", controller.FindAll)
	router.GET("/api/products/:productId", controller.FindById)
	router.POST("/api/products", controller.Create)
	router.PUT("/api/products/:productId", controller.Update)
	router.DELETE("/api/products/:productId", controller.Delete)

	router.PanicHandler = exception.ErrorHandler

	return middleware.NewAuthMiddleware(router)
}

func truncateProduct(db *sql.DB) {
	db.Exec("TRUNCATE product")
}

func TestCreateProductSuccess(t *testing.T) {
	db := SetupTestDB()
	truncateProduct(db)
	router := SetupProductRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget","price":100, "category_id":1}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/products", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}
