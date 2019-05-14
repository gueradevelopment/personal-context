package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"personal-context/db"

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

func writeResponse(w http.ResponseWriter, c chan db.Result) {
	result := <-c
	w.WriteHeader(result.Code)
	if result.Err == nil {
		marshalled, err := json.Marshal(result.Result)
		if err != nil {
			fmt.Fprintf(w, "Error!")
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(marshalled)
		}
	}
}

func writeResponseArr(w http.ResponseWriter, c chan db.ResultArray) {
	result := <-c
	w.WriteHeader(result.Code)
	if result.Err == nil {
		marshalled, err := json.Marshal(result.Result)
		if err != nil {
			fmt.Fprintf(w, "Error!")
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(marshalled)
		}
	}
}
