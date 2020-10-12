package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	fmt.Println("Web server ready ✅")

	go http.ListenAndServeTLS(
		":443",
		"/etc/letsencrypt/live/plantdex.ml/fullchain.pem",
		"/etc/letsencrypt/live/plantdex.ml/privkey.pem",
		corsMiddleware(responseMiddleware(serveSPA("dist"))),
	)
	http.ListenAndServe(
		":80",
		corsMiddleware(responseMiddleware(http.HandlerFunc(redirectToHTTPS))),
	)
}

func serveSPA(directory string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := filepath.Join(directory, filepath.Clean(r.URL.Path))

		if _, err := os.Stat(p); err != nil {
			http.ServeFile(w, r, filepath.Join(directory, "index.html"))
		} else {
			http.ServeFile(w, r, p)
		}
	}
}

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	domain := strings.Split(r.Host, ":")[0]

	http.Redirect(
		w,
		r,
		"https://"+domain+r.URL.Path,
		http.StatusMovedPermanently,
	)
}

func responseMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "go")

		h.ServeHTTP(w, r)
	}
}

func corsMiddleware(h http.Handler) http.HandlerFunc {
	apiURL := os.Getenv("API_URL")

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Security-Policy", "default-src 'self' "+apiURL)
		w.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Add(
			"Strict-Transport-Security",
			"max-age=63072000; includeSubDomains; preload",
		)
		w.Header().Add("X-Content-Type-Options", "nosniff")
		w.Header().Add("X-Frame-Options", "SAMEORIGIN")
		w.Header().Add("X-XSS-Protection", "1; mode=block")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	}
}
