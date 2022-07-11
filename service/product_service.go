package service

import (
	"context"
	"golang-rest-api/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	SellProduct(ctx context.Context, request web.ProductSellRequest) web.ProductResponse
	Delete(ctx context.Context, ProductId int)
	FindById(ctx context.Context, ProductId int) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
}
