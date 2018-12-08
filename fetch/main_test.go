package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/tjcain/ukpolice"
)

func TestUpdateRequired(t *testing.T) {
	ds := mockDataStore{}
	client := ukpolice.NewClient(&http.Client{})

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	url, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	mux.HandleFunc("/crimes-street-dates", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
		[
			{
				"date" : "2018-06",
				"stop-and-search" : [
					"bedfordshire",
					"cleveland",
					"durham"
				]
			}
		]`)
	})

	client.BaseURL = url
	got := UpdateRequired(ds, client)
	want := true
	if got != want {
		t.Errorf("UpdateRequired() returned %v want %v", got, want)
	}
}

type mockDataStore struct{}

func (mds mockDataStore) GetDate() string {
	return "2018-01"
}
