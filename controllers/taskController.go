package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gueradevelopment/personal-context/models"

	"github.com/gorilla/mux"
	"github.com/gueradevelopment/personal-context/db"
)

var (
	data db.TaskDB
)

func get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := make(chan db.Result)
	go data.Get(id, c)
	result := <-c
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

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c := make(chan db.ResultArray)
	go data.GetAll(c)
	result := <-c
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

func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := make(chan db.Result)
	go data.Delete(id, c)
	result := <-c
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

func edit(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item models.Task
	json.Unmarshal(body, &item)

	c := make(chan db.Result)
	go data.Edit(item, c)
	result := <-c
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

func add(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item models.Task
	json.Unmarshal(body, &item)

	c := make(chan db.Result)
	go data.Add(item, c)
	result := <-c
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

// AddTaskController function
func AddTaskController(r *mux.Router) {
	data = db.TaskDB{}
	r.HandleFunc("/", getAll).Methods(http.MethodGet)
	r.HandleFunc("/{id}", get).Methods(http.MethodGet)
	r.HandleFunc("/", edit).Methods(http.MethodPut)
	r.HandleFunc("/", add).Methods(http.MethodPost)
	r.HandleFunc("/{id}", delete).Methods(http.MethodDelete)

}
