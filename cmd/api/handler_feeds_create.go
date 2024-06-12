package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func (api ApiHandler) HandlerFeedsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feed, err := api.DbStore.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	api.RespondWithJson(w, http.StatusOK, feed)
}
