package stats

// Stat describes statistical information returned form the database.
type Stat struct {
	Name  string `json:"name" db:"name"`
	Count int    `json:"count" db:"count"`
}
