package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	var newProd Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProd)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Give me valid json", 400)
		return
	}

	newProd.ID = len(productList) + 1
	productList = append(productList, newProd)

	sendData(w, newProd, 201)
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // CORS
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json") // Response as JSON
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux() // mux = router

	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("OPTIONS /products", http.HandlerFunc(getProducts))

	mux.HandleFunc("POST /create-products", http.HandlerFunc(createProduct))
	mux.HandleFunc("OPTIONS /create-products", http.HandlerFunc(createProduct))

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

Advanced Routing
----------------


*/
