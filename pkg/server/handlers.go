package server

import "net/http"

// GetSchema
func (s *Server) GetSchemaOverView(w http.ResponseWriter, r *http.Request) {
	// check boltDB for cached "schemaOverview"

}
