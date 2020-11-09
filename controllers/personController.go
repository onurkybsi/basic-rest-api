package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/basic-rest-api/db"
	"github.com/basic-rest-api/models"
	"github.com/gorilla/mux"
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
	r.Body.Close()

	res := db.GetPeopleByName(searchModel.Name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// InsertPerson : Insert given Person object
func InsertPerson(w http.ResponseWriter, r *http.Request) {
	var insertedPerson models.Person

	err := json.NewDecoder(r.Body).Decode(&insertedPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	r.Body.Close()

	res := db.InsertPerson(insertedPerson)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// UpdatePerson : Update given Person object by Id
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var updatedPerson models.Person

	err := json.NewDecoder(r.Body).Decode(&updatedPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	r.Body.Close()

	res := db.UpdatePersonByID(vars["ID"], updatedPerson)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// DeletePerson : Delete Person by Id
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	res := db.DeletePersonByID(vars["ID"])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
