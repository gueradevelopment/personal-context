package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gueradevelopment/personal-context/routers"
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
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		panic(err)
	}
}
