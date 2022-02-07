package migrations

import "gorm.io/gorm"

type Role struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"size:255;index:idx_name,unique"`
	User User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	gorm.Model
}
