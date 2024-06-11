package app

import (
	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/cmd/api"
	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/cmd/database"
)

type App struct {
	Api api.ApiHandler
}

func NewApp() App {
	db := database.NewDb()
	api := api.NewApi(db)
	return App{
		Api: api,
	}
}
