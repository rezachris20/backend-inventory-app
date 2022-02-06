package migrations

import "gorm.io/gorm"

type Transaction struct {
	ID                uint `gorm:"primaryKey"`
	Diskon            int
	TotalBayar        int
	UserID            int
	TransactionDetail TransactionDetail
	gorm.Model
}

type TransactionDetail struct {
	ID            uint `gorm:"primaryKey"`
	TransactionID int
	ProductID     int
	Qty           int
	HargaSatuan   int
	Total         int
	gorm.Model
}
