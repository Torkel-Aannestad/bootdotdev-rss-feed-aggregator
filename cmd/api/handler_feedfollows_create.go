package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func (api ApiHandler) HandlerFeedfollowsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Feed uuid.UUID `json:"feed"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feedFollow, err := api.DbStore.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{

		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.Feed,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	api.RespondWithJson(w, http.StatusOK, feedFollow)
}
