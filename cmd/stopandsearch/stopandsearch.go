package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tjcain/ukpolice"

	"github.com/tjcain/stopandsearch/pkg/fetch"
	"github.com/tjcain/stopandsearch/pkg/http/rest"
	"github.com/tjcain/stopandsearch/pkg/stats"
	"github.com/tjcain/stopandsearch/pkg/storage/postgres"
)

func main() {
	// wait for containers to spool up properly
	for i := 0; i < 5; i++ {
		// set up storage
		_, err := postgres.NewPostgresDB()
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	db, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatalln("Db setup failed:", err)
	}
	defer db.Close()

	// set up services.
	customClient := http.Client{Timeout: time.Second * 120}
	client := ukpolice.NewClient(&customClient)
	fetch := fetch.NewService(db, client)

	// create tables
	fetch.CreateTables()

	// if err := fetch.UpdateData(); err != nil {
	// 	log.Fatalf("UpdateData Failed: %s", err)
	// }

	// set up services.
	stats := stats.NewService(db)

	// set up router
	r := rest.Handler(stats)
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatalln(err)
	}

}
