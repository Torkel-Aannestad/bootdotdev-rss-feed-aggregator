package api

import (
	"fmt"
	"net/http"

	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/internal/auth"
)

func (api ApiHandler) UsersRetrieveByAPIKey(w http.ResponseWriter, r *http.Request) {

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "no auth header")
		return
	}
	fmt.Printf("apiKey: %v\n", apiKey)
	user, err := api.DbStore.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "No user to be found")
		return
	}

	api.RespondWithJson(w, http.StatusOK, api.DatabaseUserToUser(user))
}
