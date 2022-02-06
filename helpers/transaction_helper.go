package helpers

import (
	"backend-inventory-app/entity"
	"backend-inventory-app/web/transaction"
)

func ToTransactionResponse(transactionEntity entity.Transaction) transaction.TransactionResponses {
	return transaction.TransactionResponses{
		ID:         transactionEntity.ID,
		Diskon:     transactionEntity.Diskon,
		TotalBayar: transactionEntity.TotalBayar,
		User:       transactionEntity.User,
		CreatedAt:  transactionEntity.CreatedAt,
		UpdatedAt:  transactionEntity.UpdatedAt,
	}
}

func ToTransactionResponses(transactions []entity.Transaction) []transaction.TransactionResponses {
	var transactionResponses []transaction.TransactionResponses

	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, ToTransactionResponse(transaction))
	}

	return transactionResponses
}
