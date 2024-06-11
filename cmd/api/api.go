package api

import "github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/cmd/database"

type ApiHandler struct {
	DbStore database.DbClient
}

func NewApi(db database.DbClient) ApiHandler {
	return ApiHandler{
		DbStore: db,
	}
}
