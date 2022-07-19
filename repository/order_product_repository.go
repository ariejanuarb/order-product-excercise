package repository

import (
	"context"
	"database/sql"
	"golang-rest-api/model/domain"
)

type OrdersProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, op domain.OrdersProduct) domain.OrdersProduct
	Update(ctx context.Context, tx *sql.Tx, op domain.OrdersProduct) domain.OrdersProduct
	Delete(ctx context.Context, tx *sql.Tx, op domain.OrdersProduct)
	FindById(ctx context.Context, tx *sql.Tx, opId int) (domain.OrdersProduct, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.OrdersProduct
}
