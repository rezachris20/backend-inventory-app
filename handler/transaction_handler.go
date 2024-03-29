package handler

import (
	"backend-inventory-app/service"
	"backend-inventory-app/web"
	"backend-inventory-app/web/transaction"
	userPayload "backend-inventory-app/web/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionHandler interface {
	CreateTransaction(c echo.Context) (err error)
	FindAllTransactions(c echo.Context) (err error)
	FindAllTransactionsByUserID(c echo.Context) (err error)
}

type transactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(service service.TransactionService) TransactionHandler {
	return &transactionHandler{service: service}
}

func (h *transactionHandler) CreateTransaction(c echo.Context) (err error) {
	input := new(transaction.CreateTransactionRequest)

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

	saveTransaction, err := h.service.CreateTransaction(*input, userID)
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
		Data:   saveTransaction,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) FindAllTransactions(c echo.Context) (err error) {
	transactions, err := h.service.GetAllTransactions()
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "success",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   transactions,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) FindAllTransactionsByUserID(c echo.Context) (err error) {
	user := c.Get("currentUser").(userPayload.UserResponse)
	userID := user.ID

	transactionsByUserID, err := h.service.GetAllTransactionsByUserID(userID)
	if err != nil {
		response := web.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "success",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   transactionsByUserID,
	}
	return c.JSON(http.StatusOK, response)
}
