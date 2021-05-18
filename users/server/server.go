package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
)

type Server struct {
	svr *http.Server
}

func New(
	createSvc creater,
	authenticateSvc authenticater,
	authorizeSvc authorizer,
	generateSvc generater,
	revokeSvc revoker,
	infoSvc infoer,
	domain string,
) *Server {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Group(func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				h := NewHandler(createSvc, authenticateSvc, authorizeSvc, generateSvc, revokeSvc, infoSvc, domain)

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
	fmt.Println("Users server ready ✅")

	err := s.svr.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	fmt.Println("Stopping server...")

	return s.svr.Shutdown(ctx)
}
