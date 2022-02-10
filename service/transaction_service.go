package service

import (
	"backend-inventory-app/entity"
	"backend-inventory-app/helpers"
	"backend-inventory-app/repository"
	"backend-inventory-app/web/transaction"
	"errors"
)

type TransactionService interface {
	CreateTransaction(request transaction.CreateTransactionRequest, userIDCreate int) (transaction.TransactionResponses, error)
	GetAllTransactions() ([]transaction.TransactionResponses, error)
	GetAllTransactionsByUserID(userID int) ([]transaction.TransactionResponses, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
	productRepository     repository.ProductRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository, productRepository repository.ProductRepository) TransactionService {
	return &transactionService{transactionRepository, productRepository}
}

func (s *transactionService) CreateTransaction(request transaction.CreateTransactionRequest, userIDCreate int) (transaction.TransactionResponses, error) {

	var detailTransactions []entity.TransactionDetail
	var transactionResponse transaction.TransactionResponses

	var transaction entity.Transaction

	transaction.Diskon = request.Diskon
	transaction.TotalBayar = request.TotalBayar
	transaction.UserID = userIDCreate

	for _, detail := range request.DetailTransactions {

		product, err := s.productRepository.Product(detail.ProductID)

		if err != nil || product.ID == 0 {
			return transactionResponse, errors.New("product not found")
		}

		data := entity.TransactionDetail{
			ProductID:   detail.ProductID,
			Qty:         detail.Qty,
			HargaSatuan: detail.HargaSatuan,
			Total:       detail.Total,
		}

		detailTransactions = append(detailTransactions, data)
	}
	transaction.TransactionDetail = detailTransactions

	saveMain, err := s.transactionRepository.CreateMain(transaction)
	if err != nil {
		return transactionResponse, errors.New("failed to create transaction")
	}

	return helpers.ToTransactionResponse(saveMain), nil

}

func (s *transactionService) GetAllTransactions() ([]transaction.TransactionResponses, error) {
	transactions, err := s.transactionRepository.Transactions()
	if err != nil {
		return helpers.ToTransactionResponses(transactions), err
	}

	return helpers.ToTransactionResponses(transactions), nil
}

func (s *transactionService) GetAllTransactionsByUserID(userID int) ([]transaction.TransactionResponses, error) {
	transactions, err := s.transactionRepository.TransactionsByUserID(userID)
	if err != nil {
		return helpers.ToTransactionResponses(transactions), err
	}

	return helpers.ToTransactionResponses(transactions), nil
}
