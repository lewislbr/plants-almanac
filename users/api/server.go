package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	u "users/user"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

// Server defines the server struct.
var Server = &http.Server{}

// Start initializes the server.
func Start(cr u.CreateService, an u.AuthenticateService, az u.AuthorizeService, gn u.GenerateService) error {
	port := os.Getenv("USERS_PORT")
	Server.Addr = ":" + port

	router := httprouter.New()
	handler := NewHandler(cr, an, az, gn)

	router.HandlerFunc("POST", "/signup", handler.CreateUser)
	router.HandlerFunc("POST", "/login", handler.LogInUser)
	router.HandlerFunc("GET", "/authorize", handler.AuthorizeUser)
	router.HandlerFunc("GET", "/refresh", handler.RefreshToken)

	Server.Handler = corsMiddleware(router)

	Server.IdleTimeout = 120 * time.Second
	Server.ReadTimeout = 5 * time.Second
	Server.WriteTimeout = 10 * time.Second

	fmt.Println("Users server ready ✅")

	err := Server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return errors.Wrap(err, "")
	}

	return nil
}

// Stop stops the server.
func Stop(ctx context.Context) error {
	fmt.Println("Stopping server...")

	err := Server.Shutdown(ctx)
	if err != nil {
		return err
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
