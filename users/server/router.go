package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"users/authenticate"
	"users/authorize"
	"users/create"
	"users/generate"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

var Server = &http.Server{}

func Start(cs create.CreateService, ns authenticate.AuthenticateService, zs authorize.AuthorizeService, gs generate.GenerateService) error {
	port := os.Getenv("USERS_PORT")
	Server.Addr = ":" + port

	router := httprouter.New()
	handler := NewHandler(cs, ns, zs, gs)

	router.HandlerFunc(http.MethodPost, "/signup", handler.Create)
	router.HandlerFunc(http.MethodPost, "/login", handler.LogIn)
	router.HandlerFunc(http.MethodGet, "/authorize", handler.Authorize)
	router.HandlerFunc(http.MethodGet, "/refresh", handler.Refresh)

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
