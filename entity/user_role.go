package entity

import "time"

type UserRole struct {
	ID        int
	UserID    int
	User      User
	RoleID    int
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
