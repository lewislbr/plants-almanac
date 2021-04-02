package server

import "net/http"

func headersMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		h.ServeHTTP(w, r)
	})
}
