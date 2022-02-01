package app

import (
	"backend-inventory-app/helpers"
	"backend-inventory-app/repository"
	"backend-inventory-app/service"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(key string, c echo.Context) (bool, error) {
	db := NewDB()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	if !strings.Contains(key, "Bearer") {
		return false, nil
	}

	tokenString := ""
	arrayToken := strings.Split(key, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	//Validate token
	token, err := helpers.ValidateToken(tokenString)
	if err != nil {
		return false, nil
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, nil
	}

	userID := int(claim["user_id"].(float64))
	user, err := userService.GetUserByID(userID)
	if err != nil {
		return false, nil
	}
	c.Set("currentUser", user)

	return true, nil
}
