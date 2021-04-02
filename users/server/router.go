package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"users/authenticate"
	"users/authorize"
	"users/create"
	"users/generate"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/pkg/errors"
)

var server = &http.Server{}

func setUpRouter(h *handler, web string) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	r.Use(cors.Handler(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Origin"},
		AllowedMethods:   []string{"GET", "OPTIONS", "POST"},
		AllowedOrigins:   []string{web},
		MaxAge:           86400,
	}))
	r.Use(headersMiddleware)

	r.Post("/signup", h.Create)
	r.Post("/login", h.LogIn)
	r.Get("/authorization", h.Authorize)
	r.Get("/refresh", h.Refresh)

	return r
}

func Start(cs create.CreateService, ns authenticate.AuthenticateService, zs authorize.AuthorizeService, gs generate.GenerateService, port, web string) error {
	handler := NewHandler(cs, ns, zs, gs)
	router := setUpRouter(handler, web)

	server.Addr = ":" + port
	server.Handler = router
	server.IdleTimeout = 120 * time.Second
	server.MaxHeaderBytes = 1 << 20 // 1 MB
	server.ReadTimeout = 5 * time.Second
	server.WriteTimeout = 10 * time.Second

	fmt.Println("Users server ready ✅")

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func Stop(ctx context.Context) error {
	fmt.Println("Stopping server...")

	return server.Shutdown(ctx)
}
