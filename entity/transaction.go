package entity

import "time"

type Transaction struct {
	ID                int
	Diskon            int
	TotalBayar        int
	UserID            int
	TransactionDetail []TransactionDetail
	User              User
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type TransactionDetail struct {
	ID            int
	TransactionID int
	ProductID     int
	Qty           int
	HargaSatuan   int
	Total         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
