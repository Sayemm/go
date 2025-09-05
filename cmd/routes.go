package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /middle", manager.With(
		http.HandlerFunc(handlers.Test),
		middleware.Extra,
	))

	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
	))

	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.CreateProduct),
	))

	mux.Handle("GET /products/{productId}", manager.With(
		http.HandlerFunc(handlers.GetProductById),
	))
}
