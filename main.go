package main

import (
	"net/http"

	"personal-context/routers"
)

func main() {
	r := routers.GetRouter()
	http.ListenAndServe(":8080", r)
}
