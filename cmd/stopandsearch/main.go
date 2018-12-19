package main

import (
	"log"
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

	db.Ping()
	log.Println("Connected and pinged")

}
