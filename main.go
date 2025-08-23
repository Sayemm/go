package main

import (
	"ecommerce/cmd"
)

func main() {
	cmd.Serve()
}

/*

Advanced Routing
----------------


Middleware
==========
- GET /products -> getProducts (this function is doing handleCors and handlePreflight!! - NOT SOLID)
- GET /products -> (handleCors -> handlePreflight) Middleware -> getProducts (handler/controller)

* passing router mux to globalrouter
* so request comes to globalrouter
* globalrouter handles cors issue then preflight if it's OPTIONS
	* not OPTIONS - mux handles from there
*/
