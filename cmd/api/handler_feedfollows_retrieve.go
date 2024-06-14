package api

import (
	"net/http"

	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/internal/database"
)

func (api ApiHandler) HandlerFeedfollowsRetrieve(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollow, err := api.DbStore.ListFeedFollows(r.Context(), user.ID)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Couldn't retrieve feeds")
		return
	}
	api.RespondWithJson(w, http.StatusOK, feedFollow)
}
