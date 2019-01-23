package postgres

type Categories struct {
	AgeRange       []string `db:"age_range"`
	Ethnicity      []string `db:"ethnicity"`
	Outcome        []string `db:"outcome"`
	OutcomeLinked  []string `db:"outcome_linked_to_object"`
	ObjectOfSearch []string `db:"object_of_search"`
	Legislation    []string `db:"legislation"`
}
