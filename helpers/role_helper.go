package helpers

import (
	"backend-inventory-app/entity"
	"backend-inventory-app/web/role"
)

func ToRoleResponse(data entity.Role) role.RoleResponses {
	return role.RoleResponses{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToRoleResponses(data []entity.Role) []role.RoleResponses {
	var roleResponses []role.RoleResponses

	for _, role := range data {
		roleResponses = append(roleResponses, ToRoleResponse(role))
	}

	return roleResponses
}
