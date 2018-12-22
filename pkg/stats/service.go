package stats

import "net/url"

// Repository provides access to storage.
type Repository interface {
	// GetColumnCount returns the counts for a given column and properties query.
	GetColumnCount(column string, properties url.Values) ([]Stat, error)
}

// Service provides stat options. Currently limited to fetching stats.
type Service interface {
	GetStats(column string, properties url.Values) ([]Stat, error)
}

type service struct {
	r Repository
}

// NewService creates a stats service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetStats(column string, properties url.Values) ([]Stat, error) {
	return s.r.GetColumnCount(column, properties)
}
