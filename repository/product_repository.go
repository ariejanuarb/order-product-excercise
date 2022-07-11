package repository

import (
	"context"
	"database/sql"
	"golang-rest-api/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product
	SellProduct(ctx context.Context, tx *sql.Tx, productId int, qty int) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, Product domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, ProductId int) (domain.Product, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.Product
}
