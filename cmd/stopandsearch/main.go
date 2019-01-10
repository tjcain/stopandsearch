package main

import (
	"net/http"

	"github.com/tjcain/stopandsearch/pkg/http/rest"
	"github.com/tjcain/stopandsearch/pkg/stats"
	"github.com/tjcain/stopandsearch/pkg/storage/postgres"

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
	db := postgres.NewPostgresDB()
	defer db.Close()

	// set up services.
	stats := stats.NewService(db)

	// set up router
	r := rest.Handler(stats)
	http.ListenAndServe(":4000", r)

}
