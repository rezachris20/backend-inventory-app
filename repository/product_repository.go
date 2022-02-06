package repository

import (
	"backend-inventory-app/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product entity.Product) (entity.Product, error)
	Products() ([]entity.Product, error)
	Product(productID int) (entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	Delete(productID int) (bool, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product entity.Product) (entity.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) Products() ([]entity.Product, error) {
	var products []entity.Product

	err := r.db.Preload("Category").Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *productRepository) Product(productID int) (entity.Product, error) {

	var product entity.Product
	err := r.db.Preload("Category").Where("id = ?", productID).Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) Update(product entity.Product) (entity.Product, error) {
	err := r.db.Omit("Category").Updates(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) Delete(productID int) (bool, error) {

	var product entity.Product
	err := r.db.Omit("Category").Where("id = ?", productID).Delete(&product).Error
	if err != nil {
		return false, nil
	}

	return true, nil
}
