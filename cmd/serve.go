package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() // mux = router

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
