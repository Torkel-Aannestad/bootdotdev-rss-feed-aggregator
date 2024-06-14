package api

import (
	"net/http"

	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func (api ApiHandler) HandlerFeedfollowsDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	id := r.URL.Query().Get("feedFollowID")

	feedFollowID, err := uuid.Parse(id)
	if err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "Couldn't parse feedFollowID")
		return
	}

	err = api.DbStore.DeleteFeedFollows(r.Context(), feedFollowID)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Couldn't delete feed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
