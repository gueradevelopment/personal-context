package routers

import (
	"personal-context/controllers"
	"personal-context/middleware"

	"github.com/gorilla/mux"
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

	gueraBookController := r.PathPrefix("/guerabooks").Subrouter()
	controllers.AddGuerabookController(gueraBookController)

	return r
}
