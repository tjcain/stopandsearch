package main

import (
	"net/http"

	"github.com/tjcain/stopandsearch/pkg/http/rest"
	"github.com/tjcain/stopandsearch/pkg/stats"
	"github.com/tjcain/stopandsearch/pkg/storage/postgres"
)

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
