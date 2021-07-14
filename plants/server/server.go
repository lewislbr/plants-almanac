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
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

type Server struct {
	svr *http.Server
}

func New(plantSvc plantService, authUrl, webUrl string) *Server {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	r.Use(cors.Handler(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Origin"},
		AllowedMethods:   []string{"DELETE", "GET", "OPTIONS", "PUT", "POST"},
		AllowedOrigins:   []string{webUrl},
		MaxAge:           86400,
	}))
	r.Use(headersMiddleware, authzMiddleware(authUrl))

	h := NewHandler(plantSvc)

	r.Post("/", h.Add)
	r.Get("/", h.ListAll)
	r.Get("/{id}", h.ListOne)
	r.Put("/{id}", h.Edit)
	r.Delete("/{id}", h.Delete)

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
