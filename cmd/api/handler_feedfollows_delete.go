package api

import (
	"encoding/json"
	"net/http"

	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func (api ApiHandler) HandlerFeedfollowsDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	err = api.DbStore.DeleteFeedFollows(r.Context(), params.FeedID)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Couldn't delete feed")
		return
	}

	api.RespondWithJson(w, http.StatusNoContent, nil)
}
