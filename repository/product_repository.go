package repository

import (
	"database/sql"
	"golang-rest-api/model/domain"
)

type ProductRepository interface {
	Save(sql *sql.DB, product domain.Product) domain.Product
	Search(productName string) []domain.Product
}
