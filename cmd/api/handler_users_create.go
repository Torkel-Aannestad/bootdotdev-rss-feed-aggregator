package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func (api ApiHandler) HandlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := api.DbStore.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	api.RespondWithJson(w, http.StatusOK, api.DatabaseUserToUser(user))
}
