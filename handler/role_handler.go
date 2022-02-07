package handler

import (
	"backend-inventory-app/service"
	"backend-inventory-app/web"
	"backend-inventory-app/web/role"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoleHandler interface {
	CreateRole(c echo.Context) (err error)
	EditRole(c echo.Context) (err error)
	DeleteRole(c echo.Context) (err error)
	GetRole(c echo.Context) (err error)
	GetRoles(c echo.Context) (err error)
}

type roleHandler struct {
	service service.RoleService
}

func NewRoleHandler(service service.RoleService) RoleHandler {
	return &roleHandler{service}
}

func (h *roleHandler) CreateRole(c echo.Context) (err error) {
	input := new(role.CreateOrUpdateRoleRequest)

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

	role, err := h.service.CreateNewRole(*input)
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
		Data:   role,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *roleHandler) EditRole(c echo.Context) (err error) {
	input := new(role.CreateOrUpdateRoleRequest)

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

	roleID, _ := strconv.Atoi(c.Param("id"))

	role, err := h.service.FindByID(roleID)
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	update, err := h.service.UpdateRole(role.ID, *input)
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := web.ApiResponse{
		Code:   http.StatusBadRequest,
		Status: "error",
		Data:   update,
	}
	return c.JSON(http.StatusBadRequest, response)
}

func (h *roleHandler) DeleteRole(c echo.Context) (err error) {
	roleID, _ := strconv.Atoi(c.Param("id"))

	role, err := h.service.FindByID(roleID)
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	_, err = h.service.DeleteRole(role.ID)
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
		Message: "success deleted role",
	}
	return c.JSON(http.StatusOK, response)
}

func (h *roleHandler) GetRole(c echo.Context) (err error) {
	roleID, _ := strconv.Atoi(c.Param("id"))

	role, err := h.service.FindByID(roleID)
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
		Data:   role,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *roleHandler) GetRoles(c echo.Context) (err error) {
	roles, err := h.service.FindAll()
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
		Data:   roles,
	}
	return c.JSON(http.StatusOK, response)
}
