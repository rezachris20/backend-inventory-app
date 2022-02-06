package helpers

import (
	"backend-inventory-app/entity"
	"backend-inventory-app/web/product"
	"os"
)

func ToProductResponse(productyEntity entity.Product) product.ProductResponse {
	return product.ProductResponse{
		ID:        productyEntity.ID,
		Name:      productyEntity.Name,
		Category:  productyEntity.Category,
		HargaBeli: productyEntity.HargaBeli,
		Qty:       productyEntity.Qty,
		Image:     os.Getenv("BASE_URL") + productyEntity.Image,
		CreatedAt: productyEntity.CreatedAt,
		UpdatedAt: productyEntity.UpdatedAt,
	}
}

func ToProductResponses(products []entity.Product) []product.ProductResponse {
	var productResponses []product.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}

	return productResponses
}
