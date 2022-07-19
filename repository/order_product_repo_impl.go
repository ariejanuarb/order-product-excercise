package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
)

type OrdersProductRepositoryImpl struct {
}

// contructor, tidak ada dependecy
func NewOrdersProductRepository() OrdersProductRepository {
	return &OrdersProductRepositoryImpl{}
}

func (c OrdersProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, item domain.OrdersProduct) domain.OrdersProduct {
	SQL := "insert into order_product(order_id, product_id, price, qty, amount) values (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, item.OrderId, item.ProductId, item.Price, item.Qty, item.Amount)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	item.Id = int(id)
	return item
}

func (c OrdersProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, item domain.OrdersProduct) domain.OrdersProduct {
	SQL := "update order_product set product_id = ?, order_id=?, price=?, qty=?, amount=? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, item.OrderId, item.ProductId, item.Price, item.Qty, item.Amount, item.Id)
	helper.PanicIfError(err)

	return item
}

func (c OrdersProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.OrdersProduct) {
	SQL := "delete from categories where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (c OrdersProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.OrdersProduct, error) {
	SQL := "select id, name from categories where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.OrdersProduct{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("categories is not found")
	}
}

func (c OrdersProductRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.OrdersProduct {
	SQL := "select id, name from categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.OrdersProduct
	for rows.Next() {
		category := domain.OrdersProduct{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
