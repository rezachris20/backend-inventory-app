package transaction

import (
	"backend-inventory-app/entity"
	"time"
)

type TransactionResponses struct {
	ID                int                      `json:"id"`
	Diskon            int                      `json:"diskon"`
	TotalBayar        int                      `json:"total_bayar"`
	User              entity.User              `json:"user_create"`
	TransactionDetail entity.TransactionDetail `json:"transaction_details"`
	CreatedAt         time.Time                `json:"created_at"`
	UpdatedAt         time.Time                `json:"updated_at"`
}
