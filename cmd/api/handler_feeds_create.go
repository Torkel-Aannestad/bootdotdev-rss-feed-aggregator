package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/cmd/types"
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
		UserID:    user.ID,
		Name:      params.Name,
		Url:       params.Url,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	_, err = api.DbStore.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		log.Println(err)
		return
	}

	api.RespondWithJson(w, http.StatusOK, types.DatabaseFeedToFeed(feed))
}
