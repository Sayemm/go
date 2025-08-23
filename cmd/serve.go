package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"log"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() // mux = router

	//*********
	controller := func(w http.ResponseWriter, r *http.Request) {
		log.Println("I am handler")
	}
	handler := http.HandlerFunc(controller)
	// mux.Handle("GET /route", handler)
	mux.Handle("GET /route", middleware.Logger(handler))
	//*********

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct)) // need to write the resouce name (REST) // route/entity will be plural
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductById))

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

*/
