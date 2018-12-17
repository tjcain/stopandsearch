package server

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
)

// DB holds a database connection.
type DB struct {
	db *sqlx.DB
}

// NewPostgresDB returns a postgres database given a provided connection string.
func NewPostgresDB(conn string) (*DB, error) {
	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

// Ping attempts to ping the database, returns an error if unsuccessful
func (db *DB) Ping() error {
	if err := db.db.Ping(); err != nil {
		return err
	}
	log.Println("DB PINGED")
	return nil
}

// Close closes the connection to the database
func (db *DB) Close() {
	db.Close()
}

// Generates a schema overview
func (db *DB) GenerateSchemaSummary() {

}
