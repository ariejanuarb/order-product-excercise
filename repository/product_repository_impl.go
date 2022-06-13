package repository

import (
	"database/sql"
	"golang-rest-api/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (p ProductRepositoryImpl) Save(sql *sql.DB, product domain.Product) domain.Product {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepositoryImpl) Search(productName string) []domain.Product {
	//TODO implement me
	panic("implement me")
}
