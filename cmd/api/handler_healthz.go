package api

import "net/http"

func (api ApiHandler) HandlerHealthz(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}
	api.RespondWithJson(w, http.StatusOK, body)
}

func (api ApiHandler) HandlerError(w http.ResponseWriter, r *http.Request) {
	api.RespondWithError(w, http.StatusOK, "Internal Server Error")
}
