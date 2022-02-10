package user_role

import (
	"backend-inventory-app/web/role"
	"backend-inventory-app/web/users"
	"time"
)

type UserRoleResponse struct {
	ID        int                `json:"id"`
	User      users.UserResponse `json:"user"`
	Role      role.RoleResponses `json:"role"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
