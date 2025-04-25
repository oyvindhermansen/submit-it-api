package forms

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/oyvindhermansen/submit-it-api/pkgs/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleCreateForm(formService Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var form Form

		err := json.NewDecoder(r.Body).Decode(&form)

		validate := validator.New()

		if err := validate.Struct(form); err != nil {
			http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		createdForm, err := formService.Create(form)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.RespondWithJSON(w, createdForm)
	}
}

func HandleList(formService Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		var (
			skip  int64 = 0
			limit int64 = 20
		)

		skipStr := query.Get("skip")
		limitStr := query.Get("limit")

		if skipStr != "" {
			s, err := strconv.ParseInt(skipStr, 10, 64)

			if err != nil {
				http.Error(w, "Invalid 'skip' parameter", http.StatusBadRequest)
				return
			}

			skip = s
		}

		if limitStr != "" {
			l, err := strconv.ParseInt(limitStr, 10, 64)

			if err != nil {
				http.Error(w, "Invalid 'limit' parameter", http.StatusBadRequest)
				return
			}

			limit = l
		}

		listOptions := ListOptions{
			Skip:  skip,
			Limit: limit,
		}

		forms, err := formService.List(listOptions)

		if err != nil {
			log.Printf("Error happened: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.RespondWithJSON(w, forms)
	}
}

func HandleFindById(formService Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		objectId, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			http.Error(w, "cannot convert string to objectId", http.StatusInternalServerError)
			return
		}

		form, err := formService.FindById(objectId)

		if err != nil {
			http.Error(w, "form not found", http.StatusNotFound)
			return
		}

		utils.RespondWithJSON(w, form)
	}
}
