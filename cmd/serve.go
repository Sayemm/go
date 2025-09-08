package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Serve() {
	cnf := config.GetConfig()

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	address := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server running on port ", address)
	err := http.ListenAndServe(address, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server: ", err)
		os.Exit(1)
	}
}
