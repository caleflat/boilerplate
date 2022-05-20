package middleware

import (
	error "github.com/caleflat/boilerplate/pkg/error"
	"net/http"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, error.New("Authorization header is missing").ToJSON(), http.StatusUnauthorized)
		}
	})
}
