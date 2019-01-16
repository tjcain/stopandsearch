package rest

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/tjcain/stopandsearch/stats"
)

// StatsResponse is the response payload for the Stats data model.
type StatsResponse struct {
	stats.Stat

	// We can add other fields to the response here
}

// NewStatsResponse generates a new stats response
func NewStatsResponse(stat stats.Stat) *StatsResponse {
	// We can add other fields to the response here
	return &StatsResponse{Stat: stat}
}

// Render satisfies the chi renderer interface
func (sr *StatsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Here we can perform pre-processing before a response is marshalled and
	// sent across the wire
	return nil
}

// StatsListResponse represents a response built from a slice of stats
type StatsListResponse []*StatsResponse

// NewStatsListResponse creates a response built from a slice of stats
func NewStatsListResponse(stats []stats.Stat) []render.Renderer {
	list := []render.Renderer{}
	for _, stat := range stats {
		list = append(list, NewStatsResponse(stat))
	}
	return list
}
