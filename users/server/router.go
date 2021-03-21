package server

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

var Server = &http.Server{}

func Start(cr u.CreateService, an u.AuthenticateService, az u.AuthorizeService, gn u.GenerateService) error {
	port := os.Getenv("USERS_PORT")
	Server.Addr = ":" + port

	router := httprouter.New()
	handler := NewHandler(cr, an, az, gn)

	router.HandlerFunc(http.MethodPost, "/signup", handler.CreateUser)
	router.HandlerFunc(http.MethodPost, "/login", handler.LogInUser)
	router.HandlerFunc(http.MethodGet, "/authorize", handler.AuthorizeUser)
	router.HandlerFunc(http.MethodGet, "/refresh", handler.RefreshToken)

	Server.Handler = corsMiddleware(router)

	Server.IdleTimeout = 120 * time.Second
	Server.ReadTimeout = 5 * time.Second
	Server.WriteTimeout = 10 * time.Second

	fmt.Println("Users server ready ✅")

	err := Server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func Stop(ctx context.Context) error {
	fmt.Println("Stopping server...")

	err := Server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
