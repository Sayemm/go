package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {
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
