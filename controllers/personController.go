package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/basic-rest-api/db"
)

// GetPeople : Get all people
func GetPeople(w http.ResponseWriter, r *http.Request) {

	res := db.GetAllPeople()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// GetPeopleByName : Get people that matches given name
func GetPeopleByName(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.Path, "/")[2]

	res := db.GetPeopleByName(name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
