package service

import (
	"backend-inventory-app/entity"
	"backend-inventory-app/helpers"
	"backend-inventory-app/repository"
	"backend-inventory-app/web/category"
	"errors"
)

type CategoryService interface {
	CreateNewCategory(request category.CreateCategoryRequest) (category.CategoryResponse, error)
	UpdateCategory(userID int, request category.UpdateCategoryRequest) (category.CategoryResponse, error)
	DeleteCategory(categoryID int) (bool, error)
	FindCategoryByID(categoryID int) (category.CategoryResponse, error)
	FindAllCategories() ([]category.CategoryResponse, error)
}

type categoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return &categoryService{repository}
}

func (s *categoryService) CreateNewCategory(request category.CreateCategoryRequest) (category.CategoryResponse, error) {

	var category entity.Category

	// apabila request.IsPrimary == false, maka request.CategoryID required
	if !request.IsPrimary {
		if request.CategoryID == 0 {
			return helpers.ToCategoryResponse(category), errors.New("category parent required")
		}

		// cek apakah category dengan id tersebut ada di dalam database
		cek, err := s.repository.Category(request.CategoryID)
		if err != nil || cek.ID == 0 {
			return helpers.ToCategoryResponse(category), errors.New("category parent not found")
		}
	}

	if !request.IsPrimary {
		if request.CategoryID != 0 {
			category.CategoryID = request.CategoryID
		}
	}

	category.Name = request.Name
	category.IsPrimary = request.IsPrimary
	category.IsActive = request.IsActive

	createCategory, err := s.repository.Save(category)
	if err != nil {
		return helpers.ToCategoryResponse(category), err
	}

	return helpers.ToCategoryResponse(createCategory), nil
}

func (s *categoryService) UpdateCategory(userID int, request category.UpdateCategoryRequest) (category.CategoryResponse, error) {

	var category entity.Category

	if request.CategoryID != 0 {
		category.CategoryID = request.CategoryID
	}

	if request.Name != "" {
		category.Name = request.Name
	}

	if request.IsPrimary {
		category.IsPrimary = request.IsPrimary
	}

	if request.IsActive {
		category.IsActive = request.IsActive
	}

	updateCategory, err := s.repository.Update(category)
	if err != nil {
		return helpers.ToCategoryResponse(category), err
	}

	return helpers.ToCategoryResponse(updateCategory), nil

}

func (s *categoryService) DeleteCategory(categoryID int) (bool, error) {
	category, err := s.repository.Category(categoryID)
	if err != nil || category.ID == 0 {
		return false, errors.New("category not found")
	}

	_, err = s.repository.Delete(category)
	if err != nil {
		return false, errors.New("Telah terjadi kesalahan")
	}

	return true, nil
}

func (s *categoryService) FindCategoryByID(categoryID int) (category.CategoryResponse, error) {
	category, err := s.repository.Category(categoryID)
	if err != nil || category.ID == 0 {
		return helpers.ToCategoryResponse(category), errors.New("category not found")
	}

	return helpers.ToCategoryResponse(category), nil
}

func (s *categoryService) FindAllCategories() ([]category.CategoryResponse, error) {
	categories, err := s.repository.Categories()
	if err != nil {
		return helpers.ToCategoryResponses(categories), err
	}

	return helpers.ToCategoryResponses(categories), nil
}
