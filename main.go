package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"personal-context/routers"

	"github.com/gorilla/handlers"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	headersOk := handlers.AllowedHeaders([]string{"X-Requested", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	r := routers.GetRouter()
	log.Printf("\nListening on %s...\n", addr)
	if err := http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(r)); err != nil {
		panic(err)
	}
}
