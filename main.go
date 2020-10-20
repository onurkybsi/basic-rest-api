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
	router.HandleFunc("/people/{name}", controllers.GetPeopleByName).Methods("GET")

	fmt.Println("Listening on 8000...")

	http.ListenAndServe(":8000", router)
}
