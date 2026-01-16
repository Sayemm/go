package cmd

import (
	"ecommerce/cart"
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	cartHandler "ecommerce/rest/handlers/cart"
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
	cartRepo := repo.NewCartRepo(dbCon)

	// domains
	userSvc := user.NewService(userRepo)
	productSvc := product.NewService(productRepo)
	cartSvc := cart.NewService(cartRepo, productRepo)

	// handlers
	productHandler := productHandler.NewHandler(middlewares, productSvc)
	userHandler := userHandler.NewHandler(cnf, userSvc)
	cartHandler := cartHandler.NewHandler(middlewares, cartSvc)

	server := rest.NewServer(cnf, productHandler, userHandler, cartHandler)
	server.Start()
}
