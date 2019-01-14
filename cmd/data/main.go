// Grab information from ukpolice database and store in our Storage solution.
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tjcain/stopandsearch/pkg/fetch"
	"github.com/tjcain/stopandsearch/pkg/storage/postgres"
	"github.com/tjcain/ukpolice"
)

func main() {
	// set up storage
	db, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatalln("could not open database connection:", err)
	}
	defer db.Close()

	// set up services.
	customClient := http.Client{Timeout: time.Second * 120}
	client := ukpolice.NewClient(&customClient)
	fetch := fetch.NewService(db, client)

	// create tables
	fetch.CreateTables()
	if err := fetch.UpdateData(); err != nil {
		log.Fatalf("UpdateData Failed: %s", err)
	}

}
