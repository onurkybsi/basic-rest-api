package main

import (
	"fmt"
	"net/http"

	"github.com/basic-rest-api/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/people", controllers.GetPeople).Methods("GET")
	router.HandleFunc("/people", controllers.GetPeopleByName).Methods("POST")
	router.HandleFunc("/people/insert", controllers.InsertPerson).Methods("POST")

	fmt.Println("Listening on 8000...")

	http.ListenAndServe(":8000", router)
}
