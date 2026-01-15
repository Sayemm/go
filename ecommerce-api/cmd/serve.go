package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	fmt.Println("Step 1: Loading configuration...")
	cnf := config.GetConfig()

	fmt.Println("Step 2: Connecting to database...")
	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dbCon.Close()

	fmt.Println("Step 3: Running database migrations...")
	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middleware.NewMiddlewares(cnf)

	// repo
	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)

	// domains
	userSvc := user.NewService(userRepo)
	productSvc := product.NewService(productRepo)

	// handlers
	productHandler := productHandler.NewHandler(middlewares, productSvc)
	userHandler := userHandler.NewHandler(cnf, userSvc)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
