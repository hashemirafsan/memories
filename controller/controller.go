package controller

import (
	"encoding/json"
	"net/http"
)

type HomeResponse struct {
	Status int
	Message string
}

func CallHome(w http.ResponseWriter, r *http.Request) {
	response := HomeResponse{ Status: http.StatusOK, Message: "Service active"}
	json.NewEncoder(w).Encode(response)
}