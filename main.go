package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oyvindhermansen/submit-it-api/pkgs/db"
	"github.com/oyvindhermansen/submit-it-api/pkgs/forms"
	"github.com/oyvindhermansen/submit-it-api/pkgs/utils"
)

func main() {

	db := db.Start()
	formRepository := forms.NewRepository(db.Collection("forms"))
	formService := forms.NewService(formRepository)

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, "Welcome to Submit.it - A form service")
	}).Methods("GET")

	router.HandleFunc("/forms", forms.HandleCreateForm(formService)).Methods("POST")
	router.HandleFunc("/forms", forms.HandleList(formService)).Methods("GET")
	router.HandleFunc("/forms/{id}", forms.HandleFindById(formService)).Methods("GET")

	http.ListenAndServe(":1337", router)

}
