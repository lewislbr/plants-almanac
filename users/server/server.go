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

type Server struct {
	srv *http.Server
}

func New(cs Creater, ns Authenticater, zs Authorizer, gs Generater, port, web string) *Server {
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

	s := &http.Server{}

	s.Addr = ":" + port
	s.Handler = r
	s.IdleTimeout = 120 * time.Second
	s.MaxHeaderBytes = 1 << 20 // 1 MB
	s.ReadTimeout = 5 * time.Second
	s.WriteTimeout = 10 * time.Second

	return &Server{
		srv: s,
	}
}

func (s *Server) Start() error {
	fmt.Println("Users server ready ✅")

	err := s.srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	fmt.Println("Stopping server...")

	return s.srv.Shutdown(ctx)
}
