package main

import (
	"log"
	"stopandsearch/pkg/storage/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	password = ""
	user     = "postgres"
	database = "stopandsearch_test"
)

func main() {
	// set up storage
	// connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
	// 	user, password, database, host, port)

	_, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatalf("Could not connect to db: %s", err)
	}

}
