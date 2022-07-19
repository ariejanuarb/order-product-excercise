package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang-rest-api/exception"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
	"golang-rest-api/model/web"
	"golang-rest-api/repository"
)

type OrdersServiceImpl struct {
	OrdersRepository        repository.OrdersRepository
	OrdersProductRepository repository.OrdersProductRepository
	DB                      *sql.DB
	Validate                *validator.Validate
}

// constructor
// dependency : repository, db, validator
func NewOrdersService(OrdersRepository repository.OrdersRepository, OrdersProductRepository repository.OrdersProductRepository, DB *sql.DB, validate *validator.Validate) OrdersService {
	return &OrdersServiceImpl{
		OrdersRepository:        OrdersRepository,
		OrdersProductRepository: OrdersProductRepository,
		DB:                      DB,
		Validate:                validate,
	}
}

func (service *OrdersServiceImpl) Create(ctx context.Context, request web.OrdersCreateRequest) web.OrdersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders := domain.Orders{
		CustomerId: request.CustomerId,
		InvoiceNo:  request.InvoiceNumber,
	}

	// header
	orders = service.OrdersRepository.Save(ctx, tx, orders)

	// loopping
	for _, detail := range request.Details {
		op := domain.OrdersProduct{
			ProductId: detail.ProductId,
			Price:     detail.Price,
			Qty:       detail.Qty,
			Amount:    detail.Amount,
			OrderId:   orders.Id,
		}
		// order_product
		op = service.OrdersProductRepository.Save(ctx, tx, op)
	}

	return helper.ToOrdersResponse(Orders)
}

func (service *OrdersServiceImpl) Update(ctx context.Context, request web.OrdersUpdateRequest) web.OrdersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Orders, err := service.OrdersRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	Orders.Name = request.Name

	Orders = service.OrdersRepository.Update(ctx, tx, Orders)

	return helper.ToOrdersResponse(Orders)
}

func (service *OrdersServiceImpl) Delete(ctx context.Context, OrdersId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Orders, err := service.OrdersRepository.FindById(ctx, tx, OrdersId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.OrdersRepository.Delete(ctx, tx, Orders)
}

func (service *OrdersServiceImpl) FindById(ctx context.Context, OrdersId int) web.OrdersResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Orders, err := service.OrdersRepository.FindById(ctx, tx, OrdersId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrdersResponse(Orders)
}

func (service *OrdersServiceImpl) FindAll(ctx context.Context) []web.OrdersResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.OrdersRepository.FindByAll(ctx, tx)

	return helper.ToOrdersResponses(categories)
}
