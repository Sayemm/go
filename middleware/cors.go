package middleware

import (
	"log"
	"net/http"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("I am CORS middlware")
		w.Header().Set("Access-Control-Allow-Origin", "*") // CORS
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json") // Response as JSON

		next.ServeHTTP(w, r)
	})
}
