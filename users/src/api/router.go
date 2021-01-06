package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

var isDevelopment = os.Getenv("MODE") == "development"

// Start initializes the REST API.
func Start() error {
	router := httprouter.New()

	router.HandlerFunc("POST", "/signup", createUser)
	router.HandlerFunc("POST", "/login", logInUser)
	router.HandlerFunc("GET", "/authorize", authorizeUser)

	fmt.Println("Users API ready ✅")

	port := os.Getenv("USERS_PORT")
	err := http.ListenAndServe(":"+port, corsMiddleware(router))
	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}

func corsMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var origin string
		if isDevelopment {
			origin = os.Getenv("WEB_DEVELOPMENT_URL")
		} else {
			origin = os.Getenv("WEB_PRODUCTION_URL")
		}

		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Add("Access-Control-Allow-Origin", origin)
		w.Header().Add("Access-Control-Max-Age", "86400")
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)

			return
		}

		h.ServeHTTP(w, r)
	}
}
