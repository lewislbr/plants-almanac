package server

import (
	"net/http"
	"os"
)

func corsMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Add("Access-Control-Allow-Origin", os.Getenv("WEB_URL"))
		w.Header().Add("Access-Control-Max-Age", "86400")
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)

			return
		}

		h.ServeHTTP(w, r)
	}
}
