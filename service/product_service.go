package service

import (
	"backend-inventory-app/entity"
	"backend-inventory-app/helpers"
	"backend-inventory-app/repository"
	"backend-inventory-app/web/product"
	"errors"
)

type ProductService interface {
	CreateProduct(request product.CreateProductRequest) (product.ProductResponse, error)
	GetAllProducts() ([]product.ProductResponse, error)
	GetProductByID(productID int) (product.ProductResponse, error)
	UpdateProduct(productID int, request product.UpdateProductRequest) (product.ProductResponse, error)
	DeleteProduct(productID int) (bool, error)
}

type productService struct {
	repositoryProduct  repository.ProductRepository
	repositoryCategory repository.CategoryRepository
}

func NewProductService(repositoryProduct repository.ProductRepository, repositoryCategory repository.CategoryRepository) ProductService {
	return &productService{repositoryProduct, repositoryCategory}
}

func (s *productService) CreateProduct(request product.CreateProductRequest) (product.ProductResponse, error) {
	var productResponse product.ProductResponse

	// cek apakah ada category tersebut
	category, err := s.repositoryCategory.Category(request.CategoryID)
	if err != nil || category.ID == 0 {
		return productResponse, errors.New("category not found")
	}

	imageURL := ""
	if request.Image != "" {
		imageURL = request.Image
	}

	data := entity.Product{
		Name:       request.Name,
		CategoryID: category.ID,
		HargaBeli:  request.HargaBeli,
		Qty:        request.Qty,
		Image:      imageURL,
	}

	create, err := s.repositoryProduct.Create(data)
	if err != nil {
		return productResponse, errors.New("failed to created product")
	}

	product, err := s.repositoryProduct.Product(create.ID)
	if err != nil {
		return productResponse, errors.New("failed to load product")
	}

	return helpers.ToProductResponse(product), nil
}

func (s *productService) GetAllProducts() ([]product.ProductResponse, error) {
	products, err := s.repositoryProduct.Products()
	if err != nil {
		return helpers.ToProductResponses(products), err
	}

	return helpers.ToProductResponses(products), nil
}

func (s *productService) GetProductByID(productID int) (product.ProductResponse, error) {
	product, err := s.repositoryProduct.Product(productID)
	if err != nil || product.ID == 0 {
		return helpers.ToProductResponse(product), errors.New("product not found")
	}

	return helpers.ToProductResponse(product), nil
}

func (s *productService) UpdateProduct(productID int, request product.UpdateProductRequest) (product.ProductResponse, error) {
	product, err := s.repositoryProduct.Product(productID)
	if err != nil || product.ID == 0 {
		return helpers.ToProductResponse(product), errors.New("product not found")
	}

	if request.CategoryID != 0 {
		// cek apakah ada category tsb d database
		category, err := s.repositoryCategory.Category(request.CategoryID)

		if err != nil || category.ID == 0 {
			return helpers.ToProductResponse(product), errors.New("category not found")
		}

		product.CategoryID = category.ID
	}
	product.Name = request.Name
	product.HargaBeli = request.HargaBeli
	product.Qty = request.Qty
	product.Image = request.Image

	updated, err := s.repositoryProduct.Update(product)
	if err != nil {
		return helpers.ToProductResponse(product), errors.New("failed updated product")
	}

	result, _ := s.repositoryProduct.Product(updated.ID)

	return helpers.ToProductResponse(result), nil
}

func (s *productService) DeleteProduct(productID int) (bool, error) {
	product, err := s.repositoryProduct.Product(productID)

	if err != nil || product.ID == 0 {
		return false, errors.New("product not found")
	}

	_, err = s.repositoryProduct.Delete(product.ID)
	if err != nil {
		return false, errors.New("delete category failed")
	}

	return true, nil
}
