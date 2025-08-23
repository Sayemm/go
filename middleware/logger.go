package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {

	/*
		controller := func(w http.ResponseWriter, r *http.Request) {

		}
		handler := http.HandlerFunc(controller)
		return handler
	*/

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println("I am middlware")
		next.ServeHTTP(w, r)

		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}

/*
INPUT - next middleware/thing that we will execute after this middleware
RETURN - what next middleware/thing expecting

--
We are creating a logger before Controller
	- mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))

	- next middleware is controller http.Handler
	- http.HandlerFunc(handlers.GetProducts) == http.handler => so return this
*/
