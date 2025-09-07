package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	mux := http.NewServeMux() // mux = route

	// CorsWithPreflight(Hudai(Logger(mux)))
	wrappedMux := manager.WrapMux(
		mux,
		middleware.Logger,
		middleware.Hudai,
		middleware.CorsWithPreflight,
	)

	InitRoutes(mux, manager)

	fmt.Println("Server running on:3000")

	// CorsWithPreflight(Hudai(Logger(mux)))
	// then mux will match routes (InitRoutes)
	// then route match then Extra(Test)
	err := http.ListenAndServe(":3000", wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server: ", err)
	}
}

// REQUEST PIPELINE => global router - hudai - logger - extra - handlers.Test
