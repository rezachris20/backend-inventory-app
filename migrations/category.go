package migrations

import (
	"gorm.io/gorm"
)

type Category struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:255;index:idx_name,unique" json:"name"`
	IsPrimary  bool
	IsActive   bool
	CategoryID int
	Product    Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	gorm.Model
}
