package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres driver
	"github.com/tjcain/stopandsearch/pkg/stats"
	"github.com/tjcain/ukpolice"
)

// Storage holds a postgres connection.
type Storage struct {
	db *sqlx.DB
}

// NewPostgresDB returns a postgres database given a provided connection string.
func NewPostgresDB() (*Storage, error) {
	db, err := sqlx.Connect("postgres", setDBConnectionInfo())
	if err != nil {
		return nil, err
	}
	return &Storage{db}, nil
}

// Close closes the db connection pool
func (s *Storage) Close() {
	s.db.Close()
}

// RETRIEVAL METHODS

// GetCount returns an int representing the result of `SELECT COUNT(*) WHERE`
func (s *Storage) GetCount(query string, values []interface{}) (int, error) {
	var count int
	q := fmt.Sprintf("SELECT count(*) FROM searches WHERE %s;", query)
	err := s.db.Get(&count, q, values...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetColumnCount returns a slice of stats representing the result of
// SELECT column, COUNT(column) ... WHERE ... GROUP BY column
func (s *Storage) GetColumnCount(column, query string, values []interface{}) ([]stats.Stat, error) {
	var st []Stat
	q := fmt.Sprintf(
		`SELECT %[1]s as name, 
		 count(%[1]s) as count 
		 FROM searches
		 WHERE %[2]s
		 GROUP BY %[1]s;`, column, query)
	err := s.db.Select(&st, q, values...)
	if err != nil {
		return nil, err
	}
	var r []stats.Stat
	for _, stat := range st {
		sStat := stats.Stat{
			Name:  stat.Name,
			Count: stat.Count,
		}
		r = append(r, sStat)
	}

	return r, nil
}

// STORAGE METHODS

// StoreSearch is responsible for storing search records into the database
func (s *Storage) StoreSearch(search ukpolice.Search) error {
	store := Search{}

	store.Force = search.Force
	store.Time = search.DateTime
	store.AgeRange = normalizeAge(search.AgeRange)
	store.Ethnicity = normalizeEthnicity(search.SelfDefinedEthnicity)
	store.Gender = search.Gender
	store.OutcomeLinkedToObject = search.OutcomeLinkedToObject
	store.ObjectOfSearch = normalizeObjectOfSearch(search.ObjectOfSearch)
	// store.Legislation = search.Legislation
	store.Outcome.SearchHappened = search.Outcome.SearchHappened
	store.Outcome.Desc = normalizeOutcomes(search.Outcome.Desc, search.Outcome.SearchHappened)

	insert := fmt.Sprintf(`
		INSERT INTO searches
		(force, time, age_range, ethnicity, search_happened, outcome,
		gender, outcome_linked_to_object, object_of_search)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`)
	_, err := s.db.Exec(insert, store.Force, store.Time, store.AgeRange,
		store.Ethnicity, store.Outcome.SearchHappened, store.Outcome.Desc,
		store.Gender, store.OutcomeLinkedToObject, store.ObjectOfSearch)
	return err
}
