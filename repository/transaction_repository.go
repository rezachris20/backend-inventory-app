package repository

import (
	"backend-inventory-app/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateMain(transaction entity.Transaction) (entity.Transaction, error)
	CreateDetails(transactionDetails []entity.TransactionDetail) (entity.Transaction, error)
	Transactions() ([]entity.Transaction, error)
	TransactionsByUserID(userID int) ([]entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) CreateDetails(transactionDetails []entity.TransactionDetail) (entity.Transaction, error) {
	var transaction entity.Transaction

	err := r.db.Create(&transactionDetails).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionRepository) CreateMain(transaction entity.Transaction) (entity.Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionRepository) Transactions() ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	if err := r.db.Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionRepository) TransactionsByUserID(userID int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	if err := r.db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}
