package routes

import (
	"github.com/hashemirafsan/memories/controller"
	"github.com/hashemirafsan/memories/middleware"
	"github.com/gorilla/mux"
)

// SetupRoutes ...
func SetupRoutes() (*mux.Router) {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.CallHome).Methods("GET")
	r.HandleFunc("/create", controller.CallCreate).Methods("POST")
	r.Use(middleware.ServiceClientMiddleware)

	return r
}