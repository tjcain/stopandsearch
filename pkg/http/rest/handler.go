package rest

// define our router here.
import (
	"encoding/json"
	"log"
	"net/http"
	"stopandsearch/pkg/stats"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Handler returns a chi router to be used in http.ListenAndServe.
func Handler(s stats.Service) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Helper -- to be replaced.
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	// RESTy routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/stats/{column}", listStats(s))
	})

	return r
}

func listStats(s stats.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Fatalln(err)
		}
		column := chi.URLParam(r, "column")
		s, err := s.GetStats(column, r.Form)
		if err != nil {
			log.Fatalln(err)
		}
		json.NewEncoder(w).Encode(s)
	}
}
