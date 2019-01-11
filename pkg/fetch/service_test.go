package fetch

// func Test_UpdateRequired_True(t *testing.T) {
// 	service, mux, teardown := setup()
// 	defer teardown()

// 	mux.HandleFunc("/crimes-street-dates", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, `[{"date":"2019-01","stop-and-search":["bedfordshire","btp"]}]`)
// 	})

// 	got, err := service.UpdateRequired()
// 	if err != nil {
// 		t.Errorf("updateRequire error: %s", err)
// 	}

// 	if got != true {
// 		t.Errorf("updateRequired: got <%t> want <%t>", got, true)
// 	}
// }

// func Test_UpdateRequired_False(t *testing.T) {
// 	service, mux, teardown := setup()
// 	defer teardown()

// 	mux.HandleFunc("/crimes-street-dates", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, `[{"date":"2015-01","stop-and-search":["bedfordshire","btp"]}]`)
// 	})

// 	got, err := service.UpdateRequired()
// 	if err != nil {
// 		t.Errorf("updateRequire error: %s", err)
// 	}

// 	if got != false {
// 		t.Errorf("updateRequired: got <%t> want <%t>", got, false)
// 	}
// }

// func setup() (s Service, mux *http.ServeMux, teardown func()) {
// 	client := ukpolice.NewClient(&http.Client{})

// 	// mux is the HTTP request multiplexer used with the test server.
// 	mux = http.NewServeMux()

// 	// server is a test HTTP server used to provide mock API responses.
// 	server := httptest.NewServer(mux)
// 	url, _ := url.Parse(server.URL)
// 	client.BaseURL = url

// 	r := &MockRepository{}
// 	s = NewService(r, client)

// 	return s, mux, server.Close
// }

// // Repository represnets a mock storage repository
// type MockRepository struct {
// 	UpdateDateInvoked bool
// }

// // GetDate invokes a mock implementation and returns a MM-YYYY string
// func (m *MockRepository) GetDate() (string, error) {
// 	return "2017-01", nil
// }

// // UpdateDate does nothing....
// func (m *MockRepository) UpdateDate(date string) error {
// 	return nil
// }

// func (m *MockRepository) StoreSearch(ukpolice.Search) error {

// 	return nil
// }
