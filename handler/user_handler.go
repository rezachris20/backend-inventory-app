package handler

import (
	"backend-inventory-app/service"
	"backend-inventory-app/web"
	userPayload "backend-inventory-app/web/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateNewUser(c echo.Context) error
	EditUser(c echo.Context) error
	FindUserByID(c echo.Context) error
	GetAllUsers(c echo.Context) error
	DeleteUser(c echo.Context) error
	Login(c echo.Context) error
}

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{service: service}
}

func (h *userHandler) CreateNewUser(c echo.Context) (err error) {
	input := new(userPayload.UserCreateRequest)

	if err = c.Bind(input); err != nil {
		response := web.ApiResponse{
			Code:   http.StatusUnprocessableEntity,
			Status: "error",
			Data:   err,
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if err = c.Validate(input); err != nil {
		response := web.ApiResponse{
			Code:   http.StatusUnprocessableEntity,
			Status: "error",
			Data:   err,
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	user, err := h.service.CreateNewUser(*input)
	if err != nil {
		response := web.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "error",
			Data:   err,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := web.ApiResponse{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   user,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *userHandler) EditUser(c echo.Context) (err error) {
	input := new(userPayload.UserUpdateRequest)

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
	user := c.Get("currentUser").(userPayload.UserResponse)
	userID := user.ID

	editUser, err := h.service.UpdateUser(userID, *input)

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
		Data:   editUser,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) FindUserByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	user, err := h.service.GetUserByID(userID)
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
		Data:   user,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetAllUsers(c echo.Context) error {
	users, err := h.service.GetUsers()
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
		Data:   users,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	_, err = h.service.DeleteUser(userID)
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := web.ApiResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "user has been deleted",
	}
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c echo.Context) (err error) {
	input := new(userPayload.UserLoginRequest)

	if err = c.Bind(input); err != nil {
		response := web.ApiResponse{
			Code:   http.StatusUnprocessableEntity,
			Status: "error",
			Data:   err,
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if err = c.Validate(input); err != nil {
		response := web.ApiResponse{
			Code:   http.StatusUnprocessableEntity,
			Status: "error",
			Data:   err,
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	user, err := h.service.LoginUser(*input)
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
		Data:   user,
	}
	return c.JSON(http.StatusOK, response)
}
