package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	mux := http.NewServeMux() // mux = router

	mux.Handle("GET /middle", manager.With(
		middleware.Hudai,
		middleware.Logger,
	)(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /route", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.Test))))
	mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", middleware.Logger(http.HandlerFunc(handlers.CreateProduct))) // need to write the resouce name (REST) // route/entity will be plural
	mux.Handle("GET /products/{productId}", middleware.Logger(http.HandlerFunc(handlers.GetProductById)))

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
