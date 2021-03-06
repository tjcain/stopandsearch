package postgres

import (
	"log"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/tjcain/stopandsearch/stats"
)

const (
	connStr = "host=localhost port=5432 user=postgres password= dbname=postgres sslmode=disable"
)

func Test_GetCount(t *testing.T) {
	store, teardown := setupDB()
	defer teardown()
	values := []interface{}{"staffordshire", "west-midlands", "18-24"}
	query := "(force = $1 OR force = $2) AND (age_range = $3)"
	got, err := store.GetCount(query, values)
	if err != nil {
		t.Fatalf("GetCountWhere failed: %s", err)
	}
	if got != 4 {
		t.Errorf("GetCount: Expected 4 got %d", got)
	}
}

func Test_GetColumnCount(t *testing.T) {
	store, teardown := setupDB()
	defer teardown()
	column := "ethnicity"
	values := []interface{}{"staffordshire", "west-midlands", "18-24"}
	query := "(force = $1 OR force = $2) AND (age_range = $3)"

	got, err := store.GetColumnCount(column, query, values)
	if err != nil {
		t.Fatalf("ColumnCount failed: %s", err)
	}
	if !reflect.DeepEqual(got, []stats.Stat{{Name: "asian", Count: 1}, {Name: "white", Count: 3}}) {
		t.Errorf("ColumnCount: Expected %v got %v", []stats.Stat{{Name: "asian", Count: 1}, {Name: "white", Count: 3}}, got)
	}
}

// DB setups

// Schema holds the SQL required to create and drop database tables
type TestSchema struct {
	create string
	drop   string
}

var defaultTestSchema = TestSchema{
	create: `
	CREATE TABLE searches_test
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
	);`,

	drop: `
	DROP TABLE IF EXISTS searches_test`,
}

func setupDB() (s *Storage, teardown func()) {
	s, err := New("postgres", "postgres", "postgres") // @TODO: Remove hard coding
	if err != nil {
		log.Fatalln("Cannot establish db connection", err)
	}
	s.db.MustExec(defaultTestSchema.create)
	loadDefaultFixture(s.db)

	teardown = func() {
		s.db.MustExec(defaultTestSchema.drop)
		s.db.Close()
	}

	return
}

func loadDefaultFixture(db *sqlx.DB) {
	tx := db.MustBegin()
	tx.MustExec(tx.Rebind("INSERT INTO searches (force, age_range, ethnicity) VALUES (?, ?, ?)"), "staffordshire", "18-24", "white")
	tx.MustExec(tx.Rebind("INSERT INTO searches (force, age_range, ethnicity) VALUES (?, ?, ?)"), "staffordshire", "18-24", "white")
	tx.MustExec(tx.Rebind("INSERT INTO searches (force, age_range, ethnicity) VALUES (?, ?, ?)"), "staffordshire", "18-24", "white")
	tx.MustExec(tx.Rebind("INSERT INTO searches (force, age_range, ethnicity) VALUES (?, ?, ?)"), "staffordshire", "10-17", "black")
	tx.MustExec(tx.Rebind("INSERT INTO searches (force, age_range, ethnicity) VALUES (?, ?, ?)"), "west-midlands", "10-17", "black")
	tx.MustExec(tx.Rebind("INSERT INTO searches (force, age_range, ethnicity) VALUES (?, ?, ?)"), "west-midlands", "18-24", "asian")
	tx.Commit()
}
