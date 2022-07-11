package repository

import (
	"context"
	"database/sql"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (p ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "insert into product(name, price, categoryId) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.CategoryId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	return product
}

func (p ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepositoryImpl) SellProduct(ctx context.Context, tx *sql.Tx, productId int, qty int) domain.Product {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, Product domain.Product) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, ProductId int) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	//TODO implement me
	panic("implement me")
}
