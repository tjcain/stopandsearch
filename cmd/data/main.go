// Grab information from ukpolice database and store in our Storage solution.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tjcain/stopandsearch/pkg/fetch"
	"github.com/tjcain/stopandsearch/pkg/storage/postgres"
	"github.com/tjcain/ukpolice"
)

func main() {
	// set up storage
	db := postgres.NewPostgresDB()
	defer db.Close()

	// set up services.
	customClient := http.Client{Timeout: time.Second * 120}
	client := ukpolice.NewClient(&customClient)
	fetch := fetch.NewService(db, client)

	// create tables
	fetch.CreateTables()

	fmt.Println("starting....")
	start := time.Now()
	if err := fetch.UpdateData(); err != nil {
		log.Fatalf("UpdateData Failed: %s", err)
	}

	done := time.Since(start)
	fmt.Printf("Done after %s seconds", done)

}
