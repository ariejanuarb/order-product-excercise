package controller

import (
	"github.com/julienschmidt/httprouter"
	"golang-rest-api/helper"
	"golang-rest-api/model/web"
	"golang-rest-api/service"
	"net/http"
	"strconv"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

// contructor new products controller
// dependency : products service
func NewProductController(ProductService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: ProductService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ProductCreateRequest := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(request, &ProductCreateRequest)

	ProductResponse := controller.ProductService.Create(request.Context(), ProductCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ProductUpdateRequest := web.ProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &ProductUpdateRequest)

	ProductId := params.ByName("ProductId")
	id, err := strconv.Atoi(ProductId)
	helper.PanicIfError(err)

	ProductUpdateRequest.Id = id

	ProductResponse := controller.ProductService.Update(request.Context(), ProductUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ProductId := params.ByName("ProductId")
	id, err := strconv.Atoi(ProductId)
	helper.PanicIfError(err)

	controller.ProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ProductId := params.ByName("ProductId")
	id, err := strconv.Atoi(ProductId)
	helper.PanicIfError(err)

	ProductResponse := controller.ProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ProductResponses := controller.ProductService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
