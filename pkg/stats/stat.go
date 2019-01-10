package stats

// Stat describes statistical information returned form the database.
type Stat struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
