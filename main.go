package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "My name is Sayemmm")
}

func main() {
	mux := http.NewServeMux() // mux = router

	mux.HandleFunc("/hello", helloHandler) // /hello - route
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running on:3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the server: ", err)
	}

}

/*
- We can create multiple route using router
- request goes to router and router knows which request will go to which route
- then the function that is registered with that router will be executed

*/
