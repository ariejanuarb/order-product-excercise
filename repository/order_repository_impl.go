package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
)

type OrdersRepositoryImpl struct {
}

// contructor, tidak ada dependecy
func NewOrdersRepository() OrdersRepository {
	return &OrdersRepositoryImpl{}
}

func (c OrdersRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders {
	SQL := "insert into orders(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, orders.CustomerId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	orders.Id = int(id)
	return orders
}

func (c OrdersRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders {
	SQL := "update categories set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orders.CustomerId, orders.Id)
	helper.PanicIfError(err)

	return orders
}

func (c OrdersRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, orders domain.Orders) {
	SQL := "delete from categories where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orders.Id)
	helper.PanicIfError(err)
}

func (c OrdersRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, OrdersId int) (domain.Orders, error) {
	SQL := "select id, name from categories where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, OrdersId)
	helper.PanicIfError(err)
	defer rows.Close()

	Orders := domain.Orders{}
	if rows.Next() {
		err := rows.Scan(&Orders.Id, &Orders.CustomerId)
		helper.PanicIfError(err)
		return Orders, nil
	} else {
		return Orders, errors.New("categories is not found")
	}
}

func (c OrdersRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.Orders {
	SQL := "select id, name from categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Orders
	for rows.Next() {
		Orders := domain.Orders{}
		err := rows.Scan(&Orders.Id, &Orders.CustomerId)
		helper.PanicIfError(err)
		categories = append(categories, Orders)
	}
	return categories
}
