package api

import "github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/internal/database"

type ApiHandler struct {
	DbStore *database.Queries
}

func NewApi(db *database.Queries) ApiHandler {
	return ApiHandler{
		DbStore: db,
	}
}
