package handler

import (
	"backend-inventory-app/helpers"
	"backend-inventory-app/service"
	"backend-inventory-app/web"
	"backend-inventory-app/web/product"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ProductHandler interface {
	CreateNewProduct(c echo.Context) (err error)
	GetProducts(c echo.Context) (err error)
	GetProduct(c echo.Context) (err error)
	EditProduct(c echo.Context) (err error)
	DeleteProduct(c echo.Context) (err error)
}

type productHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) ProductHandler {
	return &productHandler{service}
}

var now = time.Now()

func (h *productHandler) CreateNewProduct(c echo.Context) (err error) {
	input := new(product.CreateProductRequest)

	if err = c.Bind(input); err != nil {
		response := web.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if err = c.Validate(input); err != nil {
		response := web.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	// Upload Image
	file, err := c.FormFile("image")
	path := ""

	if err != nil {
		path = ""
	} else {
		sec := now.Unix()
		fileName := strconv.FormatInt(sec, 10)
		path = "images/products/" + fileName + ".jpg"
		helpers.UploadFile(file, path)
	}
	// Upload Image

	input.Image = path

	createProduct, err := h.service.CreateProduct(*input)
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := web.ApiResponse{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   createProduct,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *productHandler) GetProducts(c echo.Context) (err error) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   products,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetProduct(c echo.Context) (err error) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	product, err := h.service.GetProductByID(productID)
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusNotFound,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusNotFound, response)
	}

	response := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   product,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *productHandler) EditProduct(c echo.Context) (err error) {
	input := new(product.UpdateProductRequest)

	if err = c.Bind(input); err != nil {
		response := web.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	update, err := h.service.UpdateProduct(productID, *input)
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   update,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *productHandler) DeleteProduct(c echo.Context) (err error) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	_, err = h.service.DeleteProduct(productID)
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := web.ApiResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success deleted product",
	}
	return c.JSON(http.StatusOK, response)
}
