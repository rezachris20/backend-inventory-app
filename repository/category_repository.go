package repository

import (
	"backend-inventory-app/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(category entity.Category) (entity.Category, error)
	Update(category entity.Category) (entity.Category, error)
	Delete(category entity.Category) (bool, error)
	Categories() ([]entity.Category, error)
	Category(categoryID int) (entity.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Save(category entity.Category) (entity.Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) Update(category entity.Category) (entity.Category, error) {
	err := r.db.Save(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) Delete(category entity.Category) (bool, error) {
	err := r.db.Where("id = ? ", category.ID).Delete(&category).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *categoryRepository) Categories() ([]entity.Category, error) {
	var categories []entity.Category

	err := r.db.Find(&categories).Error
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *categoryRepository) Category(categoryID int) (entity.Category, error) {
	var category entity.Category

	err := r.db.Where("id = ?", categoryID).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
