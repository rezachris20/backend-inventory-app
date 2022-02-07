package entity

import "time"

type HakAkses struct {
	ID         int
	UserRoleID int
	UserRole   UserRole
	Module     string
	IsCreate   bool
	IsRead     bool
	IsUpdate   bool
	IsDelete   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
