package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Torkel-Aannestad/bootdotdev-rss-feed-aggregator/cmd/app"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load()

	app := app.NewApp()
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/api/users", app.Api.HandlerUsersCreate)
	mux.HandleFunc("GET /v1/api/users", app.Api.MiddlewareAuth(app.Api.HandlerUsersRetrieve))

	mux.HandleFunc("POST /v1/api/feeds", app.Api.MiddlewareAuth(app.Api.HandlerFeedsCreate))
	mux.HandleFunc("GET /v1/api/feeds", app.Api.HandlerFeedsRetrieve)

	mux.HandleFunc("GET /v1/api/healthz", app.Api.HandlerHealthz)
	mux.HandleFunc("GET /v1/api/err", app.Api.HandlerError)

	server := http.Server{
		Addr:         fmt.Sprintf(":%v", os.Getenv("PORT")),
		Handler:      mux,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}
	log.Fatal(server.ListenAndServe())
}
