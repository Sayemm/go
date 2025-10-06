package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)
	userRepo := repo.NewUserRepo()
	productRepo := repo.NewProductRepo()

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(userRepo, cnf)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
