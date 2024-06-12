package api

import (
	"net/http"

	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/internal/database"
)

func (api ApiHandler) HandlerUsersRetrieve(w http.ResponseWriter, r *http.Request, user database.User) {
	api.RespondWithJson(w, http.StatusOK, api.DatabaseUserToUser(user))
}
