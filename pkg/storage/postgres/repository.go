package postgres

import (
	"fmt"
	"log"
	"strings"

	"github.com/tjcain/ukpolice"

	db "upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

// Storage holds a postgres connection.

type Storage struct {
	db db.Database
}

// NewPostgresDB returns a postgres database given a provided connection string.
func NewPostgresDB(conn postgresql.ConnectionURL) (*Storage, error) {
	sess, err := postgresql.Open(conn)
	if err != nil {
		log.Fatalf("Error opening PGDB: %s", err)
	}
	sess.SetLogging(true) // Development only @TODO: configuration
	storage := &Storage{sess}
	return storage, nil
}

// Ping attempts to ping the database, returns an error if unsuccessful
func (s *Storage) Ping() error {
	if err := s.db.Ping(); err != nil {
		return err
	}
	log.Println("DB PINGED")
	return nil
}

// Close closes the connection to the database
func (s *Storage) Close() {
	s.db.Close()
}

// functionality below
func (s *Storage) StoreSearch(search ukpolice.Search) error {
	store := Search{}

	store.Force = search.Force
	store.Time = search.DateTime
	store.AgeRange = normalizeAge(search.AgeRange)
	store.Ethnicity = normalizeEthnicity(search.SelfDefinedEthnicity)
	store.Gender = search.Gender
	store.OutcomeLinkedToObject = search.OutcomeLinkedToObject
	store.ObjectOfSearch = search.ObjectOfSearch
	// store.Legislation = search.Legislation
	store.Outcome.SearchHappened = search.Outcome.SearchHappened
	store.Outcome.Desc = search.Outcome.Desc

	insert := fmt.Sprintf(`
		INSERT INTO wm_test_ss
		(force, month_year, time, age_range, ethnicity, search_happened, outcome,
		gender, outcome_linked_to_object, object_of_search)
		VALUES ($1 $2 $3 $4 $5 $6 $7 $8 $9 $10);`)
	_, err := s.db.Exec(insert, store.Force, store.Time, store.AgeRange,
		store.Ethnicity, store.Outcome.SearchHappened, store.Outcome,
		store.Gender, store.OutcomeLinkedToObject, store.ObjectOfSearch)

	return err
}

// GetSchemaCount @TODO: REFACTOR THIS JUNK!!!
func (s *Storage) GetSchemaCount(column string, properties map[string][]interface{}) {
	columns := []string{"age_range", "force", "month_year", "ethnicity",
		"search_happened", "outcome", "gender", "outcome_linked_to_object",
		"object_of_search"}

	var values []interface{}
	var where []string

	for _, key := range columns {
		w := []string{}
		if v, ok := properties[key]; ok {
			for _, sv := range v {
				values = append(values, sv)
				w = append(w, fmt.Sprintf("%s = $%d", key, len(values)))
			}
		}
		if len(w) != 0 {
			// Make (foo=$1 OR bar=$2).. etc
			str := fmt.Sprintf("(%s)", strings.Join(w, " OR "))
			log.Printf("STR: %s", str)
			where = append(where, str)
		}
	}

	if v, ok := properties["time"]; ok {
		if len(v) == 2 {
			for _, sv := range v {
				values = append(values, sv)
			}
			str := fmt.Sprintf("(time >= $%d AND time < $%d)", len(values)-1, len(values))
			where = append(where, str)
		}
	}

	type dbReturn struct {
		Name  string `db:"name"`
		Count int    `db:"count"`
	}

	var r []dbReturn

	// @TODO: Remove hard coding
	q := fmt.Sprintf(`
	 SELECT %s as Name,
	 COUNT(%s) as Count
	 FROM wm_test_ss
	 WHERE %s GROUP BY %s;`,
		column, column,
		strings.Join(where, " AND "), column)

	rows, err := s.db.Query(q, values...)
	if err != nil {
		log.Fatalln(err)
	}
	iter := db.NewIterator(rows)
	iter.All(&r)
}

// ========== TEMPORARY TOOLS ================

// DestructiveReset drops and remakes a table
func (s *Storage) DestructiveReset(tname string) error {
	d := drop(tname)
	_, err := s.db.Exec(d)
	if err != nil {
		return err
	}

	schema := fmt.Sprintf(`
	CREATE TABLE public.%s
	(
		force text,
		month_year text,
		time timestamp(6)
		without time zone,
		age_range text,
		ethnicity text,
		search_happened boolean,
		outcome text,
		gender text,
		outcome_linked_to_object boolean,
		object_of_search text
	)
		WITH
		(
		OIDS = FALSE
	);

		ALTER TABLE public.wm_test_ss
		OWNER to postgres;`, tname)

	_, err = s.db.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}

func drop(tname string) string {
	return `DROP TABLE IF EXISTS ` + tname
}
