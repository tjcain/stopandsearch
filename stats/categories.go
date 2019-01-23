package stats

type Categories struct {
	AgeRange       []string `json:"age_range"`
	Ethnicity      []string `json:"ethnicity"`
	Outcome        []string `json:"outcome"`
	OutcomeLinked  []string `json:"outcome_linked_to_object"`
	ObjectOfSearch []string `json:"object_of_search"`
	Legislation    []string `json:"legislation"`
}
