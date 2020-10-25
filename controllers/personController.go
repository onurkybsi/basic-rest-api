package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/basic-rest-api/db"
	"github.com/basic-rest-api/models"
)

// GetPeople : Get all people
func GetPeople(w http.ResponseWriter, r *http.Request) {

	res := db.GetAllPeople()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// GetPeopleByName : Get people that matches given name
func GetPeopleByName(w http.ResponseWriter, r *http.Request) {
	var searchModel models.SearchModel

	err := json.NewDecoder(r.Body).Decode(&searchModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := db.GetPeopleByName(searchModel.Name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
