package handler

import (
	"backend-inventory-app/service"
	"backend-inventory-app/web"
	"backend-inventory-app/web/category"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryHandler interface {
	CreateNewCategory(c echo.Context) (err error)
	UpdateCategory(c echo.Context) (err error)
	DeleteCategory(c echo.Context) (err error)
	FindCategoryByID(c echo.Context) (err error)
	FindAllCategories(c echo.Context) (err error)
}

type categoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) CategoryHandler {
	return &categoryHandler{service}
}

func (h *categoryHandler) CreateNewCategory(c echo.Context) (err error) {
	input := new(category.CreateCategoryRequest)

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

	categoryCreate, err := h.service.CreateNewCategory(*input)
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := web.ApiResponse{
		Code:    http.StatusCreated,
		Status:  "success",
		Message: "Create category success",
		Data:    categoryCreate,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *categoryHandler) UpdateCategory(c echo.Context) (err error) {
	panic("err")
}

func (h *categoryHandler) DeleteCategory(c echo.Context) (err error) {
	panic("err")
}

func (h *categoryHandler) FindCategoryByID(c echo.Context) (err error) {
	panic("err")
}

func (h *categoryHandler) FindAllCategories(c echo.Context) (err error) {
	panic("err")
}
