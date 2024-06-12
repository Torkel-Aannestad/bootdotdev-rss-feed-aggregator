package api

import (
	"net/http"

	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/internal/auth"
	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (api ApiHandler) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			api.RespondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}

		user, err := api.DbStore.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			api.RespondWithError(w, http.StatusNotFound, "Couldn't get user")
			return
		}

		handler(w, r, user)
	}
}
