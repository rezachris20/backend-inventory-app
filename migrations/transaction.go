package migrations

import "gorm.io/gorm"

type Transaction struct {
	ID                uint `gorm:"primaryKey"`
	Diskon            int
	TotalBayar        int
	UserID            int
	TransactionDetail TransactionDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
