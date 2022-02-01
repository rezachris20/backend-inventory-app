package helpers

import (
	"backend-inventory-app/entity"
	usersDomain "backend-inventory-app/web/users"
)

func ToUserResponse(user entity.User, token string) usersDomain.UserResponse {
	return usersDomain.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     token,
	}
}

func ToUserResponses(users []entity.User) []usersDomain.UserResponse {
	var userReponses []usersDomain.UserResponse
	for _, user := range users {
		userReponses = append(userReponses, ToUserResponse(user, ""))
	}
	return userReponses
}
