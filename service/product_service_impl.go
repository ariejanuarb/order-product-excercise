package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
	"golang-rest-api/model/web"
	"golang-rest-api/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	item := domain.Product{
		Name:       request.Name,
		Price:      request.Price,
		CategoryId: request.CategoryId,
	}

	item = service.ProductRepository.Save(ctx, tx, item)

	return helper.ToProductResponse(item)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	//TODO implement me
	panic("implement me")
}

func (service *ProductServiceImpl) SellProduct(ctx context.Context, request web.ProductSellRequest) web.ProductResponse {
	//TODO implement me
	panic("implement me")
}

func (service *ProductServiceImpl) Delete(ctx context.Context, ProductId int) {
	//TODO implement me
	panic("implement me")
}

func (service *ProductServiceImpl) FindById(ctx context.Context, ProductId int) web.ProductResponse {
	//TODO implement me
	panic("implement me")
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	//TODO implement me
	panic("implement me")
}
