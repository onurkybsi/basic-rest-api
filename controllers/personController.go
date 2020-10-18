package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/basic-rest-api/db"
)

// GetPeople : Get all person
func GetPeople(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.GetAllPeople())
}
