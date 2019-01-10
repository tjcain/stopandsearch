package stats

// Repository provides access to storage.
type Repository interface {
	// GetCount returns an int representing the result of `SELECT COUNT(*) WHERE`
	GetCount(query string, values []interface{}) (int, error)
	// GetColumnCount returns a slice of stats representing the result of
	// SELECT column, COUNT(column) ... WHERE ... GROUP BY column
	GetColumnCount(column, query string, values []interface{}) ([]Stat, error)
}

// Service provides stat options.
type Service interface {
	GetCount(query string, values []interface{}) (int, error)
	GetColumnCount(column, query string, values []interface{}) ([]Stat, error)
}

type service struct {
	r Repository
}

// NewService creates a stats service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetCount(query string, values []interface{}) (int, error) {
	return s.r.GetCount(query, values)
}

func (s *service) GetColumnCount(column, query string, values []interface{}) ([]Stat, error) {
	return s.r.GetColumnCount(column, query, values)
}
