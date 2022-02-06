package migrations

import "gorm.io/gorm"

type Product struct {
	ID                uint   `gorm:"primaryKey"`
	Name              string `gorm:"size:255;index:idx_name,unique" json:"name"`
	CategoryID        int
	HargaBeli         uint
	Qty               int
	Image             string
	TransactionDetail TransactionDetail
	gorm.Model
}
