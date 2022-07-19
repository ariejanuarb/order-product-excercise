package service

import (
	"context"
	"golang-rest-api/model/web"
)

type OrdersService interface {
	Create(ctx context.Context, request web.OrdersCreateRequest) web.OrdersResponse
	Update(ctx context.Context, request web.OrdersUpdateRequest) web.OrdersResponse
	Delete(ctx context.Context, OrdersId int)
	FindById(ctx context.Context, OrdersId int) web.OrdersResponse
	FindAll(ctx context.Context) []web.OrdersResponse
}
