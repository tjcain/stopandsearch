package postgres

import (
	"fmt"
	"log"
	"os"
)

// setDBConnectionInfo returns a connection string for postgres based on env
// variables, or if they are unset returns default settings.
func setDBConnectionInfo() string {
	host := os.Getenv("PG_HOST")
	// port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DBNAME")

	psqlInfo := ""

	if host != "" {
		psqlInfo = fmt.Sprintf("host=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, user, password, dbname)
		log.Println("CONNECTION INFO", psqlInfo)
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			"localhost", "5432", "postgres", "", "postgres")
	}

	return psqlInfo
}

// CreateTables will drop any existing tables and create new. Panics on fail.
func (s *Storage) CreateTables() {
	s.db.MustExec("DROP TABLE IF EXISTS searches")
	s.db.MustExec(create)
}

const create = `
	CREATE TABLE searches
	(
		force text,
		time timestamp(6)
		without time zone,
		age_range text,
		ethnicity text,
		search_happened boolean,
		outcome text,
		gender text,
		outcome_linked_to_object boolean,
		object_of_search text
	);`
