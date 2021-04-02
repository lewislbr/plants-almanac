package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/pkg/errors"
)

var server = &http.Server{}

func New(cs Creater, ns Authenticater, zs Authorizer, gs Generater, port, web string) *http.Server {
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

	h := NewHandler(cs, ns, zs, gs)

	r.Post("/signup", h.Create)
	r.Post("/login", h.LogIn)
	r.Get("/authorization", h.Authorize)
	r.Get("/refresh", h.Refresh)

	server.Addr = ":" + port
	server.Handler = r
	server.IdleTimeout = 120 * time.Second
	server.MaxHeaderBytes = 1 << 20 // 1 MB
	server.ReadTimeout = 5 * time.Second
	server.WriteTimeout = 10 * time.Second

	return server
}

func Start() error {
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
