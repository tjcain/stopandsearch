package postgres

type Overview struct {
	AgeRange map[string]int `json:"age_range,omitempty"`
}
