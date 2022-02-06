package migrations

import (
	"gorm.io/gorm"
)

type Category struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:255;index:idx_name,unique" json:"name"`
	IsPrimary  bool   `json:"is_primary"`
	IsActive   bool   `json:"is_active"`
	CategoryID int    `json:"category_id"`
	Product    Product
	gorm.Model
}
