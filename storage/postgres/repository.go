package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres driver
	"github.com/pkg/errors"
	"github.com/tjcain/stopandsearch/stats"
	"github.com/tjcain/ukpolice"
)

// Storage holds a postgres connection.
type Storage struct {
	db *sqlx.DB
}

// New returns a PostgreSQL back Storage.
func New(username, password, dbname, host string) (*Storage, error) {
	dsn := fmt.Sprintf("sslmode=disable dbname=%s user=%s password=%s host=%s",
		dbname, username, password, host)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "count not open connection to postgres")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "count not ping postgres")
	}

	return &Storage{db}, nil
}

// Close closes the db connection pool
func (s *Storage) Close() {
	s.db.Close()
}

// RETRIEVAL METHODS

// GetCategories fetches DISTINCT rows for each category
func (s *Storage) GetCategories() (*stats.Categories, error) {
	c := Categories{}

	q := fmt.Sprintf("SELECT DISTINCT age_range FROM searches;")
	err := s.db.Select(&c.AgeRange, q)
	if err != nil {
		errors.Wrap(err, "could not get category: age_range")
		return nil, err
	}

	q = fmt.Sprintf("SELECT DISTINCT ethnicity FROM searches;")
	err = s.db.Select(&c.Ethnicity, q)
	if err != nil {
		errors.Wrap(err, "could not get category: ethnicity")
		return nil, err
	}

	q = fmt.Sprintf("SELECT DISTINCT outcome FROM searches;")
	err = s.db.Select(&c.Outcome, q)
	if err != nil {
		errors.Wrap(err, "could not get category: outcome")
		return nil, err
	}

	q = fmt.Sprintf("SELECT DISTINCT outcome_linked_to_object FROM searches;")
	err = s.db.Select(&c.OutcomeLinked, q)
	if err != nil {
		errors.Wrap(err, "could not get category: outcome_linked_to_object")
		return nil, err
	}

	q = fmt.Sprintf("SELECT DISTINCT object_of_search FROM searches;")
	err = s.db.Select(&c.ObjectOfSearch, q)
	if err != nil {
		errors.Wrap(err, "could not get category: object_of_search")
		return nil, err
	}

	q = fmt.Sprintf("SELECT DISTINCT legislation FROM searches;")
	err = s.db.Select(&c.Legislation, q)
	if err != nil {
		errors.Wrap(err, "could not get category: legislation")
		return nil, err
	}

	stats := &stats.Categories{
		AgeRange:       c.AgeRange,
		Ethnicity:      c.Ethnicity,
		Outcome:        c.Outcome,
		OutcomeLinked:  c.OutcomeLinked,
		ObjectOfSearch: c.ObjectOfSearch,
		Legislation:    c.Legislation,
	}
	return stats, nil
}

// GetCount returns an int representing the result of `SELECT COUNT(*) WHERE`
func (s *Storage) GetCount(query string, values []interface{}) (int, error) {
	var count int
	if len(query) == 0 {
		q := fmt.Sprintf("SELECT count(*) FROM searches")
		err := s.db.Get(&count, q)
		if err != nil {
			return 0, err
		}
	}

	if len(query) > 0 {
		q := fmt.Sprintf("SELECT count(*) FROM searches WHERE %s;", query)
		err := s.db.Get(&count, q, values...)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

// GetColumnCount returns a slice of stats representing the result of
// SELECT column, COUNT(column) ... WHERE ... GROUP BY column
func (s *Storage) GetColumnCount(column, query string, values []interface{}) ([]stats.Stat, error) {
	var st []Stat
	if len(query) == 0 {
		q := fmt.Sprintf(
			`SELECT %[1]s as name, 
			 count(%[1]s) as count 
			 FROM searches
			 GROUP BY %[1]s;`, column)
		err := s.db.Select(&st, q)
		if err != nil {
			return nil, err
		}
	}

	if len(query) > 0 {
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
	store.Legislation = normalizeLegislation(search.Legislation)
	store.Outcome.SearchHappened = search.Outcome.SearchHappened
	store.Outcome.Desc = normalizeOutcomes(search.Outcome.Desc, search.Outcome.SearchHappened)

	insert := fmt.Sprintf(`
		INSERT INTO searches
		(force, time, age_range, ethnicity, search_happened, outcome,
		gender, outcome_linked_to_object, object_of_search, legislation)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`)
	_, err := s.db.Exec(insert, store.Force, store.Time, store.AgeRange,
		store.Ethnicity, store.Outcome.SearchHappened, store.Outcome.Desc,
		store.Gender, store.OutcomeLinkedToObject, store.ObjectOfSearch,
		store.Legislation)
	return err
}
