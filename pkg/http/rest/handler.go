package rest

// define our router here.
import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/tjcain/stopandsearch/pkg/stats"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// Handler returns a chi router to be used in http.ListenAndServe.
func Handler(s stats.Service) chi.Router {
	r := chi.NewRouter()

	// Basic CORS
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)

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
	r.Route("/api/stats", func(r chi.Router) {
		r.Use(ParseQueryParams)
		r.Get("/columns/{column}", columnStats(s))
		r.Get("/count", count(s))
	})

	return r
}

func columnStats(s stats.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		column := chi.URLParam(r, "column")
		query, values := getParamsFromContext(r.Context())
		log.Println("QUERY:", query)
		s, err := s.GetColumnCount(column, query, values)
		if err != nil {
			log.Fatalln(err)
		}
		json.NewEncoder(w).Encode(s)
	}
}

func count(s stats.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, values := getParamsFromContext(r.Context())
		log.Println(query)
		s, err := s.GetCount(query, values)
		if err != nil {
			log.Fatalln(err)
		}
		json.NewEncoder(w).Encode(s)
	}

}

func getParamsFromContext(ctx context.Context) (string, []interface{}) {
	c := ctx.Value(keyQueryParams)

	// type assertion
	query := c.(struct {
		Query  string
		Values []interface{}
	}).Query

	// type assertion
	values := c.(struct {
		Query  string
		Values []interface{}
	}).Values

	return query, values
}
