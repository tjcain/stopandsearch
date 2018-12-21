package overview

// Repository provides access to storage
type Repository interface {
	// GetSchemaCount returns counts of elements within the schema
	GetSchemaCount() Overview
}

// Service provides overview listing operations.
type Service interface {
	ShowSchema() Overview
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

// ShowSchema returns an overview of schema statistics.
func (s *service) ShowSchema() Overview {
	return s.r.GetSchemaCount()
}
