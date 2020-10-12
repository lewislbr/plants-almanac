package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	httpsServer := newHTTPSServer(
		corsMiddleware(responseMiddleware(serveSPA("dist"))),
	)

	redirectToHTTPS()

	fmt.Println("Web server ready ✅")

	httpsServer.ListenAndServeTLS(
		"/etc/letsencrypt/live/plantdex.ml/fullchain.pem",
		"/etc/letsencrypt/live/plantdex.ml/privkey.pem",
	)
}

func newHTTPSServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         ":443",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig: &tls.Config{
			NextProtos:       []string{"h2"},
			MinVersion:       tls.VersionTLS13,
			CurvePreferences: []tls.CurveID{tls.CurveP256, tls.X25519},
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
			PreferServerCipherSuites: true,
		},
	}
}

func redirectToHTTPS() {
	redirectServer := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Connection", "close")
			url := "https://" + r.Host + r.URL.String()
			http.Redirect(w, r, url, http.StatusMovedPermanently)
		}),
	}

	go func() {
		log.Fatal(redirectServer.ListenAndServe())
	}()
}

func serveSPA(directory string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file := filepath.Join(directory, filepath.Clean(r.URL.Path))
		_, err := os.Stat(file)
		acceptedEncodings := r.Header.Get("Accept-Encoding")
		brotliExtension := ".br"

		if filepath.Clean(r.URL.Path) == "/" {
			file = file + "/index.html"
		}
		if os.IsNotExist(err) {
			file = filepath.Join(directory, "index.html")
		}
		if strings.Contains(acceptedEncodings, "br") {
			w.Header().Add("Content-Encoding", "br")
			w.Header().Add("Vary", "Accept-Encoding")

			switch {
			case strings.Contains(file, ".html"):
				w.Header().Add("Content-Type", "text/html")
			case strings.Contains(file, ".css"):
				w.Header().Add("Content-Type", "text/css")
			case strings.Contains(file, ".js"):
				w.Header().Add("Content-Type", "application/javascript")
			case strings.Contains(file, ".svg"):
				w.Header().Add("Content-Type", "image/svg+xml")
			}

			http.ServeFile(w, r, file+brotliExtension)
		} else {
			http.ServeFile(w, r, file)
		}
	}
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
