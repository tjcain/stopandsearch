package postgres

// Stat describes statistical information returned form the database.
type Stat struct {
	Name  string `db:"name"`
	Count int    `db:"count"`
}
