package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Controller interface to wrap controllers
type Controller interface {
	Get(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Edit(http.ResponseWriter, *http.Request)
	Add(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	AddController(*mux.Router)
}
