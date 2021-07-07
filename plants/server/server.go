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

func New(plantSvc plantService, authUrl string) *Server {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Group(func(r chi.Router) {
		r.Use(authzMiddleware(authUrl))
		r.Route("/api", func(r chi.Router) {
			r.Route("/plants", func(r chi.Router) {
				h := NewHandler(plantSvc)

				r.Post("/", h.Add)
				r.Get("/", h.ListAll)
				r.Get("/{id}", h.ListOne)
				r.Put("/{id}", h.Edit)
				r.Delete("/{id}", h.Delete)
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
	log.Println("Plants server ready ✅")

	err := s.svr.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("error starting server: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	log.Println("Stopping server...")

	err := s.svr.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("error stopping server: %w", err)
	}

	return nil
}
