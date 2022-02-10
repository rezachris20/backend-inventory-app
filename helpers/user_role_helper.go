package helpers

import (
	"backend-inventory-app/entity"
	"backend-inventory-app/web/user_role"
)

func ToUserRoleResponse(data entity.UserRole) user_role.UserRoleResponse {
	return user_role.UserRoleResponse{
		ID:        data.ID,
		User:      ToUserResponse(data.User, ""),
		Role:      ToRoleResponse(data.Role),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToUserRoleResponses(data []entity.UserRole) []user_role.UserRoleResponse {
	var userRoleResponses []user_role.UserRoleResponse

	for _, userRole := range data {
		userRoleResponses = append(userRoleResponses, ToUserRoleResponse(userRole))
	}

	return userRoleResponses
}
