package postgres

import (
	"os"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/tjcain/stopandsearch/pkg/stats"
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

func Test_setDBConnectionInfo_WithEnvVariables(t *testing.T) {
	setenv()
	defer unsetenv()

	want := "host=localhost port=5432 user=userName password=1234 dbname=postgres sslmode=disable"
	got := setDBConnectionInfo()

	if want != got {
		t.Errorf("DB Connection info with env variables wanted %s got %s", want, got)
	}

}

func Test_setDBConnectionInfo_default(t *testing.T) {
	want := connStr
	got := setDBConnectionInfo()

	if want != got {
		t.Errorf("DB Connection info with env variables wanted %s got %s", want, got)
	}
}

// Env setups
func setenv() {
	os.Setenv("PG_HOST", "localhost")
	os.Setenv("PG_PORT", "5432")
	os.Setenv("PG_USER", "userName")
	os.Setenv("PG_PASSWORD", "1234")
	os.Setenv("PG_DBNAME", "postgres")
}

func unsetenv() {
	os.Unsetenv("PG_HOST")
	os.Unsetenv("PG_PORT")
	os.Unsetenv("PG_USER")
	os.Unsetenv("PG_PASSWORD")
	os.Unsetenv("PG_DBNAME")
}

// DB setups

// Schema holds the SQL required to create and drop database tables
type TestSchema struct {
	create string
	drop   string
}

var defaultTestSchema = TestSchema{
	create: `
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
	);`,

	drop: `
	drop table searches`,
}

func setupDB() (s *Storage, teardown func()) {
	s = NewPostgresDB()
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
