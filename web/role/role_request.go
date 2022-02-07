package role

type CreateOrUpdateRoleRequest struct {
	Name string `json:"name" validate:"required"`
}
