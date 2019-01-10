package postgres

import (
	"fmt"
	"os"
)

// setDBConnectionInfo returns a connection string for postgres based on env
// variables, or if they are unset returns default settings.
func setDBConnectionInfo() string {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DBNAME")

	psqlInfo := ""

	if host != "" {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			"localhost", "5432", "postgres", "", "postgres")
	}

	return psqlInfo
}
