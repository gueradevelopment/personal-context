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
	taskDB db.TaskDB
)

func taskGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := make(chan db.Result)
	go taskDB.Get(id, c)
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

func taskGetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c := make(chan db.ResultArray)
	go taskDB.GetAll(c)
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

func taskDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := make(chan db.Result)
	go taskDB.Delete(id, c)
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

func taskEdit(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item models.Task
	json.Unmarshal(body, &item)

	c := make(chan db.Result)
	go taskDB.Edit(item, c)
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

func taskAdd(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item models.Task
	json.Unmarshal(body, &item)

	c := make(chan db.Result)
	go taskDB.Add(item, c)
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
	taskDB = db.TaskDB{}
	r.HandleFunc("/", taskGetAll).Methods(http.MethodGet)
	r.HandleFunc("/{id}", taskGet).Methods(http.MethodGet)
	r.HandleFunc("/", taskEdit).Methods(http.MethodPut)
	r.HandleFunc("/", taskAdd).Methods(http.MethodPost)
	r.HandleFunc("/{id}", taskDelete).Methods(http.MethodDelete)
}
