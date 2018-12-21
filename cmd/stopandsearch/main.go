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

	defer db.Close()
	v := map[string][]interface{}{
		"age_range": []interface{}{"10-17", "25-34", "under 10", "Not Stated", "foo"},
		"ethnicity": []interface{}{"White", "Black (or Black British)"},
		"time":      []interface{}{"2018-02-01", "2018-03-1"},
	}
	c := "age_range"
	db.GetSchemaCount(c, v)
	// db.SimpleQuery()
}
