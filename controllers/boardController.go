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
	boardDB db.BoardDB
)

func boardGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := make(chan db.Result)
	go boardDB.Get(id, c)
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

func boardGetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c := make(chan db.ResultArray)
	go boardDB.GetAll(c)
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

func boardDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := make(chan db.Result)
	go boardDB.Delete(id, c)
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

func boardEdit(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item models.Board
	json.Unmarshal(body, &item)

	c := make(chan db.Result)
	go boardDB.Edit(item, c)
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

func boardAdd(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item models.Board
	json.Unmarshal(body, &item)

	c := make(chan db.Result)
	go boardDB.Add(item, c)
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

// AddBoardController function
func AddBoardController(r *mux.Router) {
	boardDB = db.BoardDB{}
	r.HandleFunc("/", boardGetAll).Methods(http.MethodGet)
	r.HandleFunc("/{id}", boardGet).Methods(http.MethodGet)
	r.HandleFunc("/", boardEdit).Methods(http.MethodPut)
	r.HandleFunc("/", boardAdd).Methods(http.MethodPost)
	r.HandleFunc("/{id}", boardDelete).Methods(http.MethodDelete)
}
