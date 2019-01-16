package postgres

// CreateTables will drop any existing tables and create new. Panics on fail.
func (s *Storage) CreateTables() {
	s.db.MustExec("DROP TABLE IF EXISTS searches")
	s.db.MustExec(create)
}

const create = `
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
	);`
