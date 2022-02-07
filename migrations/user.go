package migrations

import (
	"gorm.io/gorm"
)

type User struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255"`
	Username    string `gorm:"size:255;index:idx_username,unique"`
	Email       string `gorm:"size:255;index:idx_email,unique"`
	Password    string
	Transaction Transaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserRole    UserRole    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	gorm.Model
}
