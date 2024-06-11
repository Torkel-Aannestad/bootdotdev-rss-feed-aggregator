package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api ApiHandler) RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(jsonPayload)
	w.WriteHeader(code)
}

func (api ApiHandler) RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		fmt.Printf("Internal server error: %v\n", code)
	}

	type response struct {
		Error string `json:"error"`
	}
	errorResponse := response{
		Error: msg,
	}

	api.RespondWithJson(w, code, errorResponse)
}
