package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/tjcain/ukpolice"
)

func Test_UpdateRequired_OutOfDate(t *testing.T) {
	client, mux, teardown := setup()
	ds := mockDataStore{"2016-01"}
	defer teardown()

	mux.HandleFunc("/crimes-street-dates", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"date": "2018-01"}]`)
	})

	got := UpdateRequired(ds, client)
	want := true
	if got != want {
		t.Errorf("UpdateRequired() returned %v want %v", got, want)
	}
}

func Test_UpdateRequired_InDate(t *testing.T) {
	client, mux, teardown := setup()
	ds := mockDataStore{"2018-01"}
	defer teardown()

	mux.HandleFunc("/crimes-street-dates", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"date": "2017-01"}]`)
	})

	got := UpdateRequired(ds, client)
	want := false
	if got != want {
		t.Errorf("UpdateRequired() returned %v want %v", got, want)
	}
}

func Test_UpdateRequired_SameDate(t *testing.T) {
	client, mux, teardown := setup()
	ds := mockDataStore{"2018-01"}
	defer teardown()

	mux.HandleFunc("/crimes-street-dates", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"date": "2018-01"}]`)
	})

	got := UpdateRequired(ds, client)
	want := false
	if got != want {
		t.Errorf("UpdateRequired() returned %v want %v", got, want)
	}
}

func setup() (client *ukpolice.Client, mux *http.ServeMux, teardown func()) {
	client = ukpolice.NewClient(&http.Client{})

	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	url, _ := url.Parse(server.URL) // err ignored - it is exceptionally unlikely
	client = ukpolice.NewClient(&http.Client{})
	client.BaseURL = url
	return client, mux, server.Close

}

// mockDataStore is a test data store used to provide mock database interactions.
type mockDataStore struct {
	date string
}

func (mds mockDataStore) GetDate() string {
	return mds.date
}
