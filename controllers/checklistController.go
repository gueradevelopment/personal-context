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
	checklistDB db.ChecklistDB
)

func checklistGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := make(chan db.Result)
	go checklistDB.Get(id, c)
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

func checklistGetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c := make(chan db.ResultArray)
	go checklistDB.GetAll(c)
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

func checklistDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := make(chan db.Result)
	go checklistDB.Delete(id, c)
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

func checklistEdit(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item models.Checklist
	json.Unmarshal(body, &item)

	c := make(chan db.Result)
	go checklistDB.Edit(item, c)
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

func checklistAdd(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item models.Checklist
	json.Unmarshal(body, &item)

	c := make(chan db.Result)
	go checklistDB.Add(item, c)
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

// AddChecklistController function
func AddChecklistController(r *mux.Router) {
	checklistDB = db.ChecklistDB{}
	r.HandleFunc("/", checklistGetAll).Methods(http.MethodGet)
	r.HandleFunc("/{id}", checklistGet).Methods(http.MethodGet)
	r.HandleFunc("/", checklistEdit).Methods(http.MethodPut)
	r.HandleFunc("/", checklistAdd).Methods(http.MethodPost)
	r.HandleFunc("/{id}", checklistDelete).Methods(http.MethodDelete)
}
