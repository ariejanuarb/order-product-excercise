package helper

import (
	"golang-rest-api/model/domain"
	"golang-rest-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:         product.Id,
		Name:       product.Name,
		Price:      product.Price,
		CategoryId: product.CategoryId,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var results []web.ProductResponse
	for _, item := range products {
		results = append(results, ToProductResponse(item))
	}

	return results
}
