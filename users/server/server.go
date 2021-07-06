package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	svr *http.Server
}

func New(userSvc userService, tokenSvc tokenService, domain string) *Server {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Group(func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				h := NewHandler(userSvc, tokenSvc, domain)

				r.Post("/registration", h.Create)
				r.Post("/login", h.LogIn)
				r.Post("/authorization", h.Authorize)
				r.Get("/refresh", h.Refresh)
				r.Get("/logout", h.LogOut)
				r.Get("/info", h.Info)
			})
		})
	})
	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Get("/readyz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	s := &http.Server{}

	s.Addr = ":8080"
	s.Handler = r
	s.IdleTimeout = 120 * time.Second
	s.ReadTimeout = 5 * time.Second
	s.WriteTimeout = 10 * time.Second

	return &Server{
		svr: s,
	}
}

func (s *Server) Start() error {
	log.Println("Users server ready ✅")

	err := s.svr.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("error starting server: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	log.Println("Stopping server...")

	err := s.svr.Shutdown(ctx)

	return fmt.Errorf("error stopping server: %w", err)
}
