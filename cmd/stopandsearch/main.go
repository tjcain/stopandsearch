package main

import (
	"log"
	"net/http"
	"stopandsearch/pkg/http/rest"
	"stopandsearch/pkg/stats"
	"stopandsearch/pkg/storage/postgres"

	"upper.io/db.v3/postgresql"
)

// This should come from some sort of config
var settings = postgresql.ConnectionURL{
	Database: `stopandsearch_test`,
	Host:     `localhost`,
	User:     `postgres`,
}

func main() {
	// set up storage
	db, err := postgres.NewPostgresDB(settings)
	if err != nil {
		log.Fatalf("Could not connect to db: %s", err)
	}
	defer db.Close()

	// set up services.
	stats := stats.NewService(db)

	// set up router
	r := rest.Handler(stats)
	http.ListenAndServe(":4000", r)

}
