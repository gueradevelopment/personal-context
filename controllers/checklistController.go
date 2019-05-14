package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"personal-context/models"

	"personal-context/db"

	"github.com/gorilla/mux"
)

// ChecklistController - controller for Checklist model
type ChecklistController struct {
	data db.ChecklistDB
}

// Get handler
func (controller *ChecklistController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := make(chan db.Result)
	go controller.data.Get(id, c)
	writeResponse(w, c)
}

// GetAll handler
func (controller *ChecklistController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c := make(chan db.ResultArray)
	where := r.URL.Query()
	go controller.data.GetAll(c, where)
	writeResponseArr(w, c)
}

// Delete handler
func (controller *ChecklistController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := make(chan db.Result)
	go controller.data.Delete(id, c)
	writeResponse(w, c)
}

// Edit handler
func (controller *ChecklistController) Edit(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item models.Checklist
	json.Unmarshal(body, &item)

	c := make(chan db.Result)
	go controller.data.Edit(item, c)
	writeResponse(w, c)
}

// Add handler
func (controller *ChecklistController) Add(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item models.Checklist
	json.Unmarshal(body, &item)

	c := make(chan db.Result)
	go controller.data.Add(item, c)
	writeResponse(w, c)
}

// AddController function
func (controller *ChecklistController) AddController(r *mux.Router) {
	r.HandleFunc("/", controller.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/{id}", controller.Get).Methods(http.MethodGet)
	r.HandleFunc("/", controller.Edit).Methods(http.MethodPut)
	r.HandleFunc("/", controller.Add).Methods(http.MethodPost)
	r.HandleFunc("/{id}", controller.Delete).Methods(http.MethodDelete)
}

// AddChecklistController initializer
func AddChecklistController(r *mux.Router) {
	data := db.ChecklistDB{}
	checklistController := ChecklistController{data: data}
	checklistController.AddController(r)
}
