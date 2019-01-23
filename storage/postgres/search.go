package postgres

import (
	"strings"
	"time"
)

var (
	ethnicityWhite     = "White"
	ethnicityBlack     = "Black"
	ethnicityAsian     = "Asian"
	ethnicityOther     = "Chinese or other"
	ethnicityMixed     = "Mixed"
	ethnicityNotStated = "Not stated"
)

// Search defines the Search form of a stop and search record
type Search struct {
	ID        string    `json:"id,omitempty" db:"id,omitempty"`
	Force     string    `json:"force,omitempty" db:"force,omitempty"`
	MonthYear string    `json:"month_year,omitempty" db:"month_year,omitempty"`
	Time      time.Time `json:"time,omitempty" db:"time,omitempty"`
	AgeRange  string    `json:"age_range,omitempty" db:"age_range,omitempty"`
	Ethnicity string    `json:"ethnicity,omitempty" db:"ethnicity,omitempty"`
	Outcome   struct {
		SearchHappened bool   `json:"search_happened,omitempty" db:"search_happened,omitempty"`
		Desc           string `json:"desc,omitempty" db:"outcome,omitempty"`
	} `json:"outcome,omitempty"`
	Gender                string `json:"gender,omitempty" db:"gender,omitempty"`
	OutcomeLinkedToObject bool   `json:"outcome_linked_to_object,omitempty" db:"outcome_linked_to_object,omitempty"`
	ObjectOfSearch        string `json:"object_of_search,omitempty" db:"object_of_search,omitempty"`
	Legislation           string `json:"legislation,omitempty" db:"legislation"`
}

func generateID(force, monthYear string) string {
	// @TODO: this won't work, perhaps use some sort of hash on the first n
	// characters of the entire search record?
	return strings.ToLower(force) + strings.ToLower(monthYear)
}

func normalizeEthnicity(ethnicity string) string {
	ethnicity = strings.ToLower(ethnicity)
	switch {
	case strings.HasPrefix(ethnicity, "white"):
		return ethnicityWhite
	case strings.HasPrefix(ethnicity, "black"):
		return ethnicityBlack
	case strings.HasPrefix(ethnicity, "asian"):
		return ethnicityAsian
	case strings.HasPrefix(ethnicity, "other"):
		return ethnicityOther
	case strings.HasPrefix(ethnicity, "mixed/"):
		return ethnicityMixed
	}

	return ethnicityNotStated
}

func normalizeAge(age string) string {
	if len(age) != 0 {
		return age
	}

	return "Not Stated"
}

func normalizeOutcomes(outcome string, search bool) string {
	outcome = strings.Trim(outcome, "\"")
	outcome = strings.ToLower(outcome)

	if search == false {
		return "Nothing Found"
	}
	// if search == true && outcome == "" {
	// 	return "Not Stated"
	// }

	switch {
	case strings.Contains(outcome, "arrest"):
		return "Arrest"
	case strings.Contains(outcome, "summon"):
		return "Summons"
	case strings.Contains(outcome, "penalty"):
		return "Penalty Notice"
	case strings.Contains(outcome, "caution"):
		return "Caution"
	case outcome == "khat / cannabis warning":
		return "Khat/Cannabis Warning"
	case outcome == "community resolution":
		return "Community resolution"
	default:
		return "Other"
	}
}

func normalizeObjectOfSearch(object string) string {
	switch {
	case object == "Stolen goods":
		return "Stolen property"
	case object == "Firearms":
		return "Firearms"
	case object == "Controlled drugs":
		return "Drugs"
	case object == "Psychoactive substances":
		return "Drugs"
	case object == "Articles for use in criminal damage":
		return "Criminal damage"
	case object == "Article for use in theft":
		return "Going equipped"
	case object == "Offensive weapons":
		return "Offensive weapons"
	default:
		return "Other"
	}
}

func normalizeLegislation(legislation string) string {
	switch {
	case strings.Contains(legislation, "section 60"):
		return "Section 60"
	case strings.Contains(legislation, "section 1"):
		return "PACE"
	case strings.Contains(legislation, "section 47"):
		return "Firearms Act"
	case strings.Contains(legislation, "section 23"):
		return "Misuse of Drugs Act"
	default:
		return "Other"
	}
}
