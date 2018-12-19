package postgres

import (
	"log"

	"github.com/tjcain/ukpolice"

	db "upper.io/db.v3"
	"upper.io/db.v3/postgresql"
)

// CONVERTING REPOSITORY TO UPPERIO AND THEN WRITING DATA PACKAGE TO
// FETCH STUFF FROM THE POLICE API!

// Storage holds a postgres connection.
type Storage struct {
	pg db.Database
}

// NewPostgresDB returns a postgres database given a provided connection string.
func NewPostgresDB(conn postgresql.ConnectionURL) (*Storage, error) {
	sess, err := postgresql.Open(conn)
	if err != nil {
		log.Fatalf("Error opening PGDB: %s", err)
	}
	storage := &Storage{sess}
	return storage, nil
}

// Ping attempts to ping the database, returns an error if unsuccessful
func (s *Storage) Ping() error {
	if err := s.pg.Ping(); err != nil {
		return err
	}
	log.Println("DB PINGED")
	return nil
}

// Close closes the connection to the database
func (s *Storage) Close() {
	s.pg.Close()
}

// functionality below
func (s *Storage) StoreSearch(search ukpolice.Search) error {
	store := Search{}

	store.Force = search.Force
	store.Time = search.DateTime
	store.AgeRange = search.AgeRange
	store.Ethnicity = normalizeEthnicity(search.SelfDefinedEthnicity)
	store.Gender = search.Gender
	store.OutcomeLinkedToObject = search.OutcomeLinkedToObject
	store.ObjectOfSearch = search.ObjectOfSearch
	// store.Legislation = search.Legislation
	store.Outcome.SearchHappened = search.Outcome.SearchHappened
	store.Outcome.Desc = search.Outcome.Desc

	col := s.pg.Collection("wm_test_ss")
	_, err := col.Insert(store)

	return err
}
