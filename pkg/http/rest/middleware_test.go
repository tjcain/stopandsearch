package rest

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-chi/chi"
)

func Test_ParseQueryParams(t *testing.T) {
	r := chi.NewRouter()
	r.Use(ParseQueryParams)
	r.Get("/query", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context().Value(keyQueryParams)

		// type assertion
		query := ctx.(struct {
			Query  string
			Values []interface{}
		}).Query
		queryWant := "(age_range = $1) AND (force = $2 OR force = $3)"
		if query != queryWant {
			t.Errorf("context query incorrect: got <%s> want <%s>", query, queryWant)
		}

		// type assertion
		values := ctx.(struct {
			Query  string
			Values []interface{}
		}).Values
		valuesWant := []interface{}{"18-24", "staffordshire", "west-midlands"}
		if !reflect.DeepEqual(values, valuesWant) {
			t.Errorf("context values incorrect: got <%v> want <%v>", values, valuesWant)
		}
	})

	urlQueryString := "force=staffordshire&foo=bar&force=west-midlands&age_range=18-24"

	// create test HTTP server
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+"/query?"+urlQueryString, nil)
	if err != nil {
		t.Fatalf("NewRequest failed: %s", err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("DefaultClient.Do failed: %s", err)
	}
}

func Test_ParseQueryParams_Time(t *testing.T) {
	r := chi.NewRouter()
	r.Use(ParseQueryParams)
	r.Get("/query", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context().Value(keyQueryParams)

		// type assertion
		query := ctx.(struct {
			Query  string
			Values []interface{}
		}).Query
		queryWant := "(force = $1) AND (time >= $2 AND time < $3)"
		if query != queryWant {
			t.Errorf("context query incorrect: got <%s> want <%s>", query, queryWant)
		}

		// type assertion
		values := ctx.(struct {
			Query  string
			Values []interface{}
		}).Values
		valuesWant := []interface{}{"staffordshire", "01-10-2017", "01-10-2018"}
		if !reflect.DeepEqual(values, valuesWant) {
			t.Errorf("context values incorrect: got <%v> want <%v>", values, valuesWant)
		}
	})

	urlQueryString := "force=staffordshire&time=10-2017&time=10-2018"

	// create test HTTP server
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+"/query?"+urlQueryString, nil)
	if err != nil {
		t.Fatalf("NewRequest failed: %s", err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("DefaultClient.Do failed: %s", err)
	}
}

// @TODO: Test for error values
