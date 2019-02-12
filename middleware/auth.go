package middleware

import (
	"log"
	"net/http"
)

// Auth - authentication middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TO-DO print more information about the request
		log.Println("I'll just believe you're authorized...for now")
		next.ServeHTTP(w, r)
	})
}
