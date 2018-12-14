package server

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	db *sqlx.DB
}

func NewPostgresDB(conn string) (*DB, error) {
	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}
