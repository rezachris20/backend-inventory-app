package app

import (
	"backend-inventory-app/handler"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRouter(userHandler handler.UserHandler, categoryHandler handler.CategoryHandler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	v1 := e.Group("/api/v1")

	v1.POST("/login", userHandler.Login)
	v1.POST("/register", userHandler.CreateNewUser)

	admin := v1.Group("/admin")

	admin.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:token",
		Validator: AuthMiddleware,
	}))

	userRouter := admin.Group("/user")
	userRouter.PUT("", userHandler.EditUser)
	userRouter.GET("/:id", userHandler.FindUserByID)
	userRouter.GET("", userHandler.GetAllUsers)
	userRouter.DELETE("/:id", userHandler.DeleteUser)

	categoryRouter := admin.Group("/category")
	categoryRouter.POST("", categoryHandler.CreateNewCategory)

	return e
}
