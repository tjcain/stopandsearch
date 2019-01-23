package rest

// define our router here.
import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/tjcain/stopandsearch/stats"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

var env = os.Getenv("STOP_SEARCH_ENV")

// Handler returns a chi router to be used in http.ListenAndServe.
func Handler(s stats.Service) chi.Router {
	r := chi.NewRouter()

	if env != "prod" {
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
	}

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/api/categories", getCategories(s))

	// Serve the vueSPA in history mode
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		dirPath := "/app/"
		handler := http.FileServer(http.Dir(dirPath))
		_path := r.URL.Path

		// static files
		if strings.Contains(_path, ".") || _path == "/" {
			handler.ServeHTTP(w, r)
			return
		}
		http.ServeFile(w, r, path.Join(dirPath, "/index.html"))
	})

	// RESTy routes
	r.Route("/api/stats", func(r chi.Router) {
		r.Use(middleware.URLFormat)
		r.Use(render.SetContentType(render.ContentTypeJSON))
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
		s, err := s.GetColumnCount(column, query, values)
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}
		err = render.RenderList(w, r, NewStatsListResponse(s))
		if err != nil {
			render.Render(w, r, ErrRender(err))
		}
	}
}

func count(s stats.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, values := getParamsFromContext(r.Context())
		s, err := s.GetCount(query, values)
		if err != nil {
			// @TODO: error handler
		}
		json.NewEncoder(w).Encode(s)
	}

}

func getCategories(s stats.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := s.GetCategories()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		json.NewEncoder(w).Encode(c)
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
