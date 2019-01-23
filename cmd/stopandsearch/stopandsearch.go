package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/tjcain/ukpolice"

	"github.com/tjcain/stopandsearch/config"
	"github.com/tjcain/stopandsearch/fetch"
	"github.com/tjcain/stopandsearch/http/rest"
	"github.com/tjcain/stopandsearch/stats"
	"github.com/tjcain/stopandsearch/storage/postgres"
)

var configPath = flag.String("config-file", "conf/config.json", "Path to configuration file.")

func main() {
	flag.Parse()
	c, err := config.ParseFile(*configPath)
	if err != nil {
		log.Fatalln("Error reading config:", err)
	}

	log.Println("waiting for postgres")
	time.Sleep(10 * time.Second)

	db, err := postgres.New(c.DB.Username, c.DB.Password, c.DB.Name, c.DB.Host)
	if err != nil {
		log.Fatalln("Db setup failed:", err)
	}
	defer db.Close()

	// set up services.
	client := ukpolice.NewClient(&http.Client{Timeout: time.Second * 120})
	fetch := fetch.NewService(db, client)

	// create tables
	log.Println("Starting Data Fetch")
	fetch.CreateTables()

	if err := fetch.UpdateData(); err != nil {
		log.Fatalf("UpdateData Failed: %s", err)
	}

	// set up services.
	stats := stats.NewService(db)

	// set up router
	log.Println("starting web server")
	r := rest.Handler(stats)
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatalln(err)
	}
}
