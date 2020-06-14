package routes

import (
	"github.com/hashemirafsan/logger/controller"
	"github.com/hashemirafsan/logger/middleware"
	"github.com/gorilla/mux"
)

// SetupRoutes
func SetupRoutes() (*mux.Router) {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.CallHome)
	r.Use(middleware.ServiceClientMiddleware)
	return r
}