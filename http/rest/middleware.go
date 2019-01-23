package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type key int

const (
	keyQueryParams key = iota
)

// ParseQueryParams is a middleware that parses url query parameters passed in
// as part of the url. It will generate an injection safe SQL 'WHERE' query
// string and associated slice of values and add this to the context.
// @TODO: return error if values are missing.
func ParseQueryParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// @TODO: Remove hardcoding of columns
		columns := []string{"age_range", "force", "month_year", "ethnicity",
			"search_happened", "outcome", "gender", "outcome_linked_to_object",
			"object_of_search", "legislation"}

		var values []interface{}
		var where []string

		err := r.ParseForm()
		if err != nil {
			// @TODO: Return http error of some sort
			log.Println("error logging form", err)
		}

		for _, key := range columns {
			w := []string{}
			if v, ok := r.Form[key]; ok {
				for _, sv := range v {
					values = append(values, sv)
					w = append(w, fmt.Sprintf("%s = $%d", key, len(values)))
				}
			}

			if len(w) != 0 {
				// Make (foo=$1 OR bar=$2).. etc
				str := fmt.Sprintf("(%s)", strings.Join(w, " OR "))
				where = append(where, str)
			}
		}

		// deal with time values separately
		if v, ok := r.Form["time"]; ok {
			if len(v) == 2 {
				for _, sv := range v {
					values = append(values, sv)
				}
				str := fmt.Sprintf("(time >= $%d AND time < $%d)", len(values)-1, len(values))
				where = append(where, str)
			}
		}

		q := struct {
			Query  string
			Values []interface{}
		}{
			strings.Join(where, " AND "),
			values,
		}
		ctx := context.WithValue(r.Context(), keyQueryParams, q)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
