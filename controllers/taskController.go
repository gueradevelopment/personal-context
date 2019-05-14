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

// TaskController - controller for Task model
type TaskController struct {
	data db.TaskDB
}

// Get handler
func (controller *TaskController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := make(chan db.Result)
	go controller.data.Get(id, c)
	writeResponse(w, c)
}

// GetAll handler
func (controller *TaskController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	where := r.URL.Query()
	c := make(chan db.ResultArray)
	go controller.data.GetAll(c, where)
	writeResponseArr(w, c)
}

// Delete handler
func (controller *TaskController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := make(chan db.Result)
	go controller.data.Delete(id, c)
	writeResponse(w, c)
}

// Edit handler
func (controller *TaskController) Edit(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item models.Task
	json.Unmarshal(body, &item)

	c := make(chan db.Result)
	go controller.data.Edit(item, c)
	writeResponse(w, c)
}

// Add handler
func (controller *TaskController) Add(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item models.Task
	json.Unmarshal(body, &item)

	c := make(chan db.Result)
	go controller.data.Add(item, c)
	writeResponse(w, c)
}

// AddController function
func (controller *TaskController) AddController(r *mux.Router) {
	r.HandleFunc("/", controller.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/{id}", controller.Get).Methods(http.MethodGet)
	r.HandleFunc("/", controller.Edit).Methods(http.MethodPut)
	r.HandleFunc("/", controller.Add).Methods(http.MethodPost)
	r.HandleFunc("/{id}", controller.Delete).Methods(http.MethodDelete)
}

// AddTaskController initializer
func AddTaskController(r *mux.Router) {
	data := db.TaskDB{}
	taskController := TaskController{data: data}
	taskController.AddController(r)
}
