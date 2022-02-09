package user_role

import (
	"backend-inventory-app/entity"
	"time"
)

type UserRoleResponse struct {
	ID        int         `json:"id"`
	UserID    int         `json:"user_id"`
	User      entity.User `json:"user"`
	RoleID    int         `json:"role_id"`
	Role      entity.Role `json:"role"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
