package helpers

import (
	"backend-inventory-app/entity"
	"backend-inventory-app/web/category"
)

func ToCategoryResponse(categoryEntity entity.Category) category.CategoryResponse {
	return category.CategoryResponse{
		ID:          categoryEntity.ID,
		Name:        categoryEntity.Name,
		IsPrimary:   categoryEntity.IsPrimary,
		IsActive:    categoryEntity.IsActive,
		CategoryID:  categoryEntity.CategoryID,
		SubCategory: categoryEntity.SubCategory,
		CreatedAt:   categoryEntity.CreatedAt,
		UpdatedAt:   categoryEntity.UpdatedAt,
	}
}

func ToCategoryResponses(categories []entity.Category) []category.CategoryResponse {
	var categoryResponses []category.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
