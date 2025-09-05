package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Hudai)

	mux := http.NewServeMux() // mux = router

	InitRoutes(mux, manager)

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on:3000")

	err := http.ListenAndServe(":3000", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server: ", err)
	}
}

/*
- request comes in
- route match
- execute what middleware.Logger has been returned
- start time..
- print - I am middleware
- next execute - handler (I am handlers)
- info print


In net/http, middleware is essentially a function that wraps another http.Handler.
It allows you to do something before and/or after the main handler executes.
*/
