package api

import (
	"fmt"
	"net/http"
	"os"

	u "users/src/user"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

// Start initializes the REST API.
func Start(cr u.CreateService, an u.AuthenticateService, az u.AuthorizeService) error {
	router := httprouter.New()
	handler := NewHandler(cr, an, az)

	router.HandlerFunc("POST", "/signup", handler.CreateUser)
	router.HandlerFunc("POST", "/login", handler.LogInUser)
	router.HandlerFunc("GET", "/authorize", handler.AuthorizeUser)

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
