package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Server represents the server.
type Server struct {
	Config *Config
	DB     *DB
	// other stuff
}

// New returns a new un-configured server
func New(conf *Config) *Server {
	app := &Server{Config: conf}
	return app
}

// Configure configures the server using details in Config
func (srv *Server) Configure() (err error) {
	// Currently the only service to configure is the database,
	// will add more such as boltdb, promethues, logging etc etc
	srv.DB, err = srv.Config.GetDB()
	if err != nil {
		return err
	}

	return nil

}

// Shutdown closes services within the server.
func (srv *Server) Shutdown() {
	srv.DB.Close()
	log.Println("server shutting down")
}

// NewRouter sets up routes and returns a handler
func (srv *Server) NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/schema", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("schema"))
		})
	})

	return r

}
