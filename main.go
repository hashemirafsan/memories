package main

import (
	"net/http"
	"github.com/hashemirafsan/memories/bootstrap"
	"github.com/hashemirafsan/memories/routes"
)

func main() {
	migration := bootstrap.Migration{}

	migration.Run()

	router := routes.SetupRoutes()

	http.ListenAndServe(":10000", router)
}