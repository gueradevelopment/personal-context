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

	r := routers.GetRouter()
	log.Printf("\nListening on %s...\n", addr)
	if err := http.ListenAndServe(addr, handlers.CORS()(r)); err != nil {
		panic(err)
	}
}
