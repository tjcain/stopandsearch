package main

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"net/url"
// 	"reflect"
// 	"testing"

// 	"github.com/tjcain/ukpolice"
// )

// func Test_GetUpdateSlice_OutOfDate(t *testing.T) {
// 	client, mux, teardown := setup()
// 	ds := mockDataStore{"2016-01"}
// 	defer teardown()

// 	mux.HandleFunc("/crimes-street-dates", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, `[{"date":"2018-10","stop-and-search":["bedfordshire","btp"]}]`)
// 	})

// 	got, err := GetUpdateSlice(ds, client)
// 	if err != nil {
// 		t.Errorf("GetUpdateSlice() returned error: %s", err)
// 	}
// 	want := []ukpolice.AvailabilityInfo{
// 		{Date: "2018-10", StopAndSearch: []string{"bedfordshire", "btp"}}}
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("GetUpdateSlice() returned %v want %v", got, want)
// 	}
// }

// func Test_GetUpdateSlice_InDate(t *testing.T) {
// 	client, mux, teardown := setup()
// 	ds := mockDataStore{"2018-01"}
// 	defer teardown()

// 	mux.HandleFunc("/crimes-street-dates", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, `[{"date":"2017-10","stop-and-search":["bedfordshire","btp"]}]`)
// 	})

// 	got, err := GetUpdateSlice(ds, client)
// 	if err != nil {
// 		t.Errorf("GetUpdateSlice() returned error: %s", err)
// 	}

// 	if got != nil {
// 		t.Errorf("GetUpdateSlice() returned %v wanted nil", got)
// 	}
// }

// func Test_GetUpdateSlice_SameDate(t *testing.T) {
// 	client, mux, teardown := setup()
// 	ds := mockDataStore{"2018-01"}
// 	defer teardown()

// 	mux.HandleFunc("/crimes-street-dates", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, `[{"date":"2018-01","stop-and-search":["bedfordshire","btp"]}]]`)
// 	})

// 	got, err := GetUpdateSlice(ds, client)
// 	if err != nil {
// 		t.Errorf("GetUpdateSlice() returned error: %s", err)
// 	}
// 	if got != nil {
// 		t.Errorf("GetUpdateSlice() returned %v wanted nil", got)
// 	}
// }

// func Test_GetUpdateSlice_DateParseErrors(t *testing.T) {
// 	client, mux, teardown := setup()
// 	defer teardown()

// 	// Bad date returned from database
// 	ds := mockDataStore{"2018/01"}
// 	_, err := GetUpdateSlice(ds, client)
// 	if err == nil {
// 		t.Errorf("Expected an error")
// 	}

// 	// Bad date returned from request
// 	ds = mockDataStore{"2018-01"}
// 	mux.HandleFunc("/crimes-street-dates", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, `[{"date":"2018/01","stop-and-search":["bedfordshire","btp"]}]]`)
// 	})
// 	_, err = GetUpdateSlice(ds, client)
// 	if err == nil {
// 		t.Errorf("Expected an error")
// 	}

// }

// func Test_GetUpdateSlice_BadJSONError(t *testing.T) {
// 	client, mux, teardown := setup()
// 	defer teardown()

// 	// Bad date JSON returned from request
// 	ds := mockDataStore{"2018-01"}
// 	mux.HandleFunc("/crimes-street-dates", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, `[{"date":"2018-01","stop-and-search":["bedfordshire", btp"]}]]`)
// 	})
// 	_, err := GetUpdateSlice(ds, client)
// 	if err == nil {
// 		t.Errorf("Expected an error")
// 	}

// }

// func setup() (client *ukpolice.Client, mux *http.ServeMux, teardown func()) {
// 	client = ukpolice.NewClient(&http.Client{})

// 	// mux is the HTTP request multiplexer used with the test server.
// 	mux = http.NewServeMux()

// 	// server is a test HTTP server used to provide mock API responses.
// 	server := httptest.NewServer(mux)

// 	url, _ := url.Parse(server.URL) // err ignored - it is exceptionally unlikely
// 	client = ukpolice.NewClient(&http.Client{})
// 	client.BaseURL = url
// 	return client, mux, server.Close

// }

// // mockDataStore is a test data store used to provide mock database interactions.
// type mockDataStore struct {
// 	date string
// }

// func (mds mockDataStore) GetDate() string {
// 	return mds.date
// }
