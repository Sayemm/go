package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"fmt"
	"log"
	"net/http"
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

	fmt.Println("Step 4: Setting up routes...")
	mux := http.NewServeMux()

	// Health check endpoint (to verify server is running)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"healthy","service":"%s","version":"%s"}`,
			cnf.ServiceName, cnf.Version)
	})

	// Root endpoint (welcome message)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message":"Welcome to %s API","version":"%s"}`,
			cnf.ServiceName, cnf.Version)
	})

	// STEP 5: Start HTTP Server
	address := fmt.Sprintf(":%d", cnf.HttpPort)

	fmt.Printf("Server running on port %d\n", cnf.HttpPort)

	err = http.ListenAndServe(address, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
		os.Exit(1)
	}

	// middlewares := middleware.NewMiddlewares(cnf)

	// // repo
	// userRepo := repo.NewUserRepo(dbCon)
	// productRepo := repo.NewProductRepo(dbCon)

	// // domains
	// userSvc := user.NewService(userRepo)
	// productSvc := product.NewService(productRepo)

	// // handlers
	// productHandler := productHandler.NewHandler(middlewares, productSvc)
	// userHandler := userHandler.NewHandler(cnf, userSvc)

	// server := rest.NewServer(cnf, productHandler, userHandler)
	// server.Start()
}
