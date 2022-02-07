package migrations

import "gorm.io/gorm"

type UserRole struct {
	ID       int `gorm:"primaryKey"`
	UserID   int
	RoleID   int
	HakAkses HakAkses `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	gorm.Model
}
