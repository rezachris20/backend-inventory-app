package handler

import (
	"backend-inventory-app/service"
	"backend-inventory-app/web"
	"backend-inventory-app/web/user_role"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserRoleHandler interface {
	NewUserRole(c echo.Context) (err error)
	EditUserRole(c echo.Context) (err error)
	DeleteUserRole(c echo.Context) (err error)
	UserRole(c echo.Context) (err error)
	UserRoles(c echo.Context) (err error)
}

type userRoleHandler struct {
	userRoleService service.UserRoleService
}

func NewUserRoleHandler(userRoleService service.UserRoleService) UserRoleHandler {
	return &userRoleHandler{userRoleService: userRoleService}
}

func (h *userRoleHandler) NewUserRole(c echo.Context) (err error) {
	input := new(user_role.CreateUserRoleRequest)

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

	userRole, err := h.userRoleService.CreateNewUserRole(*input)
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
		Data:   userRole,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *userRoleHandler) EditUserRole(c echo.Context) (err error) {
	input := new(user_role.UpdateUserRoleRequest)

	userRoleID, _ := strconv.Atoi(c.Param("id"))

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

	userRole, err := h.userRoleService.UpdateUserRole(userRoleID, *input)
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
		Data:   userRole,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *userRoleHandler) DeleteUserRole(c echo.Context) (err error) {
	userRoleID, _ := strconv.Atoi(c.Param("id"))

	_, err = h.userRoleService.DeleteUserRole(userRoleID)
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
		Message: "delete user role success",
	}
	return c.JSON(http.StatusOK, response)
}

func (h *userRoleHandler) UserRole(c echo.Context) (err error) {
	userRoleID, _ := strconv.Atoi(c.Param("id"))

	userRole, err := h.userRoleService.GetUserRole(userRoleID)
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
		Data:   userRole,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *userRoleHandler) UserRoles(c echo.Context) (err error) {

	userRoles, err := h.userRoleService.GetUserRoles()
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
		Data:   userRoles,
	}
	return c.JSON(http.StatusOK, response)
}
