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
		Addr:        ":443",
		Handler:     handler,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 5 * time.Second,
		TLSConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			},
			CurvePreferences:         []tls.CurveID{tls.CurveP256, tls.X25519},
			MinVersion:               tls.VersionTLS13,
			NextProtos:               []string{"h2"},
			PreferServerCipherSuites: true,
		},
		WriteTimeout: 10 * time.Second,
	}
}

func redirectToHTTPS() {
	redirectServer := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Connection", "close")
			url := "https://" + r.Host + r.URL.String()
			http.Redirect(w, r, url, http.StatusMovedPermanently)
		}),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		log.Fatal(redirectServer.ListenAndServe())
	}()
}

func serveSPA(directory string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestedPath := filepath.Join(directory, filepath.Clean(r.URL.Path))
		acceptedEncodings := r.Header.Get("Accept-Encoding")
		brotli := "br"

		if filepath.Clean(r.URL.Path) == "/" {
			requestedPath = requestedPath + "/index.html"
		}
		if _, err := os.Stat(requestedPath); os.IsNotExist(err) {
			requestedPath = filepath.Join(directory, "index.html")
		}
		if strings.Contains(acceptedEncodings, brotli) {
			serveCompressedFile := func(mimeType string) {
				w.Header().Add("Content-Encoding", brotli)
				w.Header().Add("Content-Type", mimeType)
				http.ServeFile(w, r, requestedPath+".br")
			}

			switch filepath.Ext(requestedPath) {
			case ".html":
				serveCompressedFile("text/html")
			case ".css":
				serveCompressedFile("text/css")
			case ".js":
				serveCompressedFile("application/javascript")
			case ".svg":
				serveCompressedFile("image/svg+xml")
			default:
				http.ServeFile(w, r, requestedPath)
			}
		} else {
			http.ServeFile(w, r, requestedPath)
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
