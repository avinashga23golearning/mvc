package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/avinashga23golearning/controller"
)

func main() {
	router := mux.NewRouter()
	personController := controller.NewPersonController()
	router.HandleFunc("/person/{id}", personController.GetPersonByID).Methods("GET")
	router.HandleFunc("/person", personController.CreatePerson).Methods("POST")
	router.HandleFunc("/person/{id}", personController.DeletePersonByID).Methods("DELETE")
	router.HandleFunc("/person/{id}", personController.UpdatePerson).Methods("PUT")

	http.ListenAndServe(":8080", router)
}
