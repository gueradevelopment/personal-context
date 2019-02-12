package routers

import (
	"github.com/gorilla/mux"
	"github.com/gueradevelopment/personal-context/controllers"
	"github.com/gueradevelopment/personal-context/middleware"
)

// GetRouter function
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.Auth)
	r.Use(middleware.Logger)
	r.Use(middleware.Cors)

	taskRouter := r.PathPrefix("/tasks").Subrouter()
	controllers.AddTaskController(taskRouter)

	checklistRouter := r.PathPrefix("/checklists").Subrouter()
	controllers.AddChecklistController(checklistRouter)

	boardRouter := r.PathPrefix("/boards").Subrouter()
	controllers.AddBoardController(boardRouter)

	return r
}
