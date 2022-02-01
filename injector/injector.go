//go:build wireinject
// +build wireinject

package injector

import (
	"backend-inventory-app/app"
	"backend-inventory-app/handler"
	"backend-inventory-app/repository"
	"backend-inventory-app/service"
	"net/http"

	"github.com/google/wire"
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	handler.NewUserHandler,
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	handler.NewCategoryHandler,
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		userSet,
		categorySet,
		app.NewRouter,
		app.NewServer,
	)

	return nil
}
