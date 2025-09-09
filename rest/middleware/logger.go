package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println("I am Logger middlware")
		next.ServeHTTP(w, r)

		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}
