package repository

import (
	"context"
	"database/sql"
	"golang-rest-api/model/domain"
)

type OrdersRepository interface {
	Save(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders
	Update(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders
	Delete(ctx context.Context, tx *sql.Tx, orders domain.Orders)
	FindById(ctx context.Context, tx *sql.Tx, ordersId int) (domain.Orders, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.Orders
}
