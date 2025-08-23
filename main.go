package main

import (
	"ecommerce/cmd"
)

func main() {
	cmd.Serve()
}

/*
Middleware
==========
- GET /products -> getProducts (this function is doing handleCors and handlePreflight!! - NOT SOLID)
- GET /products -> (handleCors -> handlePreflight) Middleware -> getProducts (handler/controller)

-> GlobalRouter - Route Match - Controller
-> anything before controller is middleware
-> if we do anything after controller than controller will become middleware (no restrictions)

*/
