package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// lowercase - private property - cannot access outside main package
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // CORS
	w.Header().Set("Content-Type", "application/json") // Response as JSON

	if r.Method != http.MethodGet {
		http.Error(w, "Please give me a GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func main() {
	mux := http.NewServeMux() // mux = router

	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server running on:3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the server: ", err)
	}

}

func init() {
	pd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Photo of Orange",
		Price:       20,
		ImgUrl:      "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcQCWfkZS0Zx7Hz7LBbW1-VggJsj1vDD_2bJnBaezT3YpRDljzGkASfYN8iI3wNFCHM59cwSWAHAlBZzQmwUg_n1I-WI8loYzTib-Xs40lM",
	}

	pd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Photo of Apple",
		Price:       2,
		ImgUrl:      "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcQCWfkZS0Zx7Hz7LBbW1-VggJsj1vDD_2bJnBaezT3YpRDljzGkASfYN8iI3wNFCHM59cwSWAHAlBZzQmwUg_n1I-WI8loYzTib-Xs40lM",
	}

	pd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Photo of Banana",
		Price:       60,
		ImgUrl:      "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcQCWfkZS0Zx7Hz7LBbW1-VggJsj1vDD_2bJnBaezT3YpRDljzGkASfYN8iI3wNFCHM59cwSWAHAlBZzQmwUg_n1I-WI8loYzTib-Xs40lM",
	}

	productList = append(productList, pd1)
	productList = append(productList, pd2)
	productList = append(productList, pd3)
}

/*

GET
----
- frontend requests resource to backend and backend response as a representational state (json...) (transfer)

-> when we request to backend from frontend, we send some information
-> that information from the backend can be found in r (*http.Request)
-> when backend send the information/response, we tell w (w http.ResponseWriter) to write it

RESPONSE HAS 2 PART
-------------------
- header
- body

=====================>
The key is: in HTTP handlers, you don’t “return” data like in normal functions — instead, you write the response directly to w http.ResponseWriter.

Step-by-step
-------------

w http.ResponseWriter
	This is provided by Go’s net/http package.
	It’s like an output stream (or a pipe) connected directly to the client (frontend/browser).
	Whatever you write to w is sent back as the HTTP response.

json.NewEncoder(w)
	Creates a JSON encoder that will write directly to the HTTP response.
	You don’t store the JSON in a variable — you send it straight out.

encoder.Encode(productList)
	Takes productList, turns it into JSON, and writes that JSON into w.
	At this moment, the data is already being sent over the network to the frontend.

Why no return?
	The http.HandlerFunc type expects you to write to w, not return data.
	When your handler finishes, Go closes the response and sends it to the client.
*/
