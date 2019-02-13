package middleware

import (
	"log"
	"net/http"
)

// Logger - logging middleware
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TO-DO print more information about the request
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
