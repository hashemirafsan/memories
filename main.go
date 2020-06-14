package main

import (
	"net/http"
	"github.com/hashemirafsan/memories/routes"
)

func main() {
	r := routes.SetupRoutes()
	http.ListenAndServe(":10000", r)

}