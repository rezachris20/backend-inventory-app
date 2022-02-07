// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"backend-inventory-app/app"
	"backend-inventory-app/handler"
	"backend-inventory-app/repository"
	"backend-inventory-app/service"
	"github.com/google/wire"
	"net/http"
)

// Injectors from injector.go:

func InitializedServer() *http.Server {
	db := app.NewDB()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository, categoryRepository)
	productHandler := handler.NewProductHandler(productService)
	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository, productRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	roleRepository := repository.NewRoleRepository(db)
	roleService := service.NewRoleService(roleRepository)
	roleHandler := handler.NewRoleHandler(roleService)
	echo := app.NewRouter(userHandler, categoryHandler, productHandler, transactionHandler, roleHandler)
	server := app.NewServer(echo)
	return server
}

// injector.go:

var userSet = wire.NewSet(repository.NewUserRepository, service.NewUserService, handler.NewUserHandler)

var categorySet = wire.NewSet(repository.NewCategoryRepository, service.NewCategoryService, handler.NewCategoryHandler)

var productSet = wire.NewSet(repository.NewProductRepository, service.NewProductService, handler.NewProductHandler)

var transactionSet = wire.NewSet(repository.NewTransactionRepository, service.NewTransactionService, handler.NewTransactionHandler)

var roleSet = wire.NewSet(repository.NewRoleRepository, service.NewRoleService, handler.NewRoleHandler)
