package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"personal-context/db"
	"personal-context/models"

	"github.com/gorilla/mux"
)

// BoardController - controller for Board model
type BoardController struct {
	data db.BoardDB
}

// Get handler
func (controller *BoardController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := make(chan db.Result)
	go controller.data.Get(id, c)
	writeResponse(w, c)
}

// GetAll handler
func (controller *BoardController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c := make(chan db.ResultArray)
	where := r.URL.Query()
	go controller.data.GetAll(c, where)
	writeResponseArr(w, c)
}

// Delete handler
func (controller *BoardController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := make(chan db.Result)
	go controller.data.Delete(id, c)
	writeResponse(w, c)
}

// Edit handler
func (controller *BoardController) Edit(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item models.Board
	json.Unmarshal(body, &item)

	c := make(chan db.Result)
	go controller.data.Edit(item, c)
	writeResponse(w, c)
}

// Add handler
func (controller *BoardController) Add(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item models.Board
	json.Unmarshal(body, &item)

	c := make(chan db.Result)
	go controller.data.Add(item, c)
	writeResponse(w, c)
}

// AddController function
func (controller *BoardController) AddController(r *mux.Router) {
	r.HandleFunc("/", controller.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/{id}", controller.Get).Methods(http.MethodGet)
	r.HandleFunc("/", controller.Edit).Methods(http.MethodPut)
	r.HandleFunc("/", controller.Add).Methods(http.MethodPost)
	r.HandleFunc("/{id}", controller.Delete).Methods(http.MethodDelete)
}

// AddBoardController initializer
func AddBoardController(r *mux.Router) {
	data := db.BoardDB{}
	boardController := BoardController{data: data}
	boardController.AddController(r)
}
