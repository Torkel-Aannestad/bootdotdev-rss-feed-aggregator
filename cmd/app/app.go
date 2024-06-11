package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/cmd/api"
	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/internal/database"
)

type App struct {
	Api api.ApiHandler
}

func NewApp() App {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	dbConnection, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbClient := database.New(dbConnection)

	api := api.NewApi(dbClient)
	return App{
		Api: api,
	}
}
