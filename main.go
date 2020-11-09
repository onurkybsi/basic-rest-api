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
	router.HandleFunc("/people", controllers.GetPeopleByName).Methods("GET")
	router.HandleFunc("/people", controllers.InsertPerson).Methods("POST")
	router.HandleFunc("/people/{ID}", controllers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{ID}", controllers.DeletePerson).Methods("DELETE")

	fmt.Println("Listening on 8000...")

	http.ListenAndServe(":8000", router)
}
