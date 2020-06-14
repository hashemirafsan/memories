package controller

import (
	"github.com/hashemirafsan/memories/bootstrap"
	"github.com/hashemirafsan/memories/model"
	"encoding/json"
	"net/http"
)

// HomeResponse ...
type HomeResponse struct {
	Status int
	Message string
}

// CallHome ...
func CallHome(w http.ResponseWriter, r *http.Request) {
	response := HomeResponse{ Status: http.StatusOK, Message: "Service active"}
	json.NewEncoder(w).Encode(response)
}

// CallCreate ...
func CallCreate(w http.ResponseWriter, r *http.Request) {
	db := bootstrap.Database{}
	db.Connection().Create(&model.ActivityLog{
		ServiceClientID: 1,
		Event: "MODEL_EVENT",
		Data: "OK_FINE",
		CreatorID: "1",
		CreatorType: "Admin",
	})
	response := HomeResponse{ Status: http.StatusCreated, Message: "Event created"}
	json.NewEncoder(w).Encode(response)
}

