package user_role

type CreateUserRoleRequest struct {
	UserID int `json:"user_id" validate:"required"`
	RoleID int `json:"role_id" validate:"required"`
}

type UpdateUserRoleRequest struct {
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}
