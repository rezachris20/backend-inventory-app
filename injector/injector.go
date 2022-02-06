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
	repository.NewTransactionRepository, service.NewTransactionService, handler.NewTransactionHandler)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		userSet,
		categorySet,
		productSet,
		transactionSet,
		app.NewRouter,
		app.NewServer,
	)

	return nil
}
