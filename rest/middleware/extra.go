package middleware

import (
	"log"
	"net/http"
)

func (m *Middlewares) Extra(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("=> Extra Middleware")
		next.ServeHTTP(w, r)
	})
}
