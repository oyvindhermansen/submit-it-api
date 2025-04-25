package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

type GenericResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func RespondWithJSON(w http.ResponseWriter, payload any) {
	response, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}

	if reflect.TypeOf(payload).Kind() == reflect.String {
		payload = GenericResponse{
			Message: fmt.Sprintf("%v", payload),
			Status:  http.StatusOK,
		}

		response, err = json.Marshal(payload)

		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
