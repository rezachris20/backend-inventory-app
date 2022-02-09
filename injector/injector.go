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

var productSet = wire.NewSet(
	repository.NewProductRepository,
	service.NewProductService,
	handler.NewProductHandler,
)

var transactionSet = wire.NewSet(
	repository.NewTransactionRepository,
	service.NewTransactionService,
	handler.NewTransactionHandler,
)

var roleSet = wire.NewSet(
	repository.NewRoleRepository,
	service.NewRoleService,
	handler.NewRoleHandler,
)

var userRoleSet = wire.NewSet(
	repository.NewUserRoleRepository,
	service.NewUserRoleService,
	handler.NewUserRoleHandler,
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		userSet,
		categorySet,
		productSet,
		transactionSet,
		roleSet,
		userRoleSet,
		app.NewRouter,
		app.NewServer,
	)

	return nil
}
