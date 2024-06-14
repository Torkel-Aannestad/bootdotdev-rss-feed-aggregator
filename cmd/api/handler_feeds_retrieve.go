package api

import (
	"net/http"

	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/cmd/types"
)

func (api ApiHandler) HandlerFeedsRetrieve(w http.ResponseWriter, r *http.Request) {

	feeds, err := api.DbStore.ListFeeds(r.Context())
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Couldn't find feeds")
		return
	}

	api.RespondWithJson(w, http.StatusOK, types.DatabaseFeedsToFeeds(feeds))
}
