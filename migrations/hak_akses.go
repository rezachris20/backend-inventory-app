package migrations

import "gorm.io/gorm"

type HakAkses struct {
	ID         int `gorm:"primaryKey"`
	UserRoleID int
	Module     string
	IsCreate   bool
	IsRead     bool
	IsUpdate   bool
	IsDelete   bool
	gorm.Model
}
