package main

import (
	"net/http"

	"github.com/gueradevelopment/personal-context/routers"
)

func main() {
	r := routers.GetRouter()
	http.ListenAndServe(":8080", r)
}
