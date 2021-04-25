package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
)

type Server struct {
	srv *http.Server
}

func New(as Adder, ls Lister, es Editer, ds Deleter, auth string) *Server {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Group(func(r chi.Router) {
		r.Use(authorizationMiddleware(auth))
		r.Route("/api", func(r chi.Router) {
			r.Route("/plants", func(r chi.Router) {
				h := NewHandler(as, ls, es, ds)

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
	s.MaxHeaderBytes = 1 << 20 // 1 MB
	s.ReadTimeout = 5 * time.Second
	s.WriteTimeout = 10 * time.Second

	return &Server{
		srv: s,
	}
}

func (s *Server) Start() error {
	fmt.Println("Plants server ready ✅")

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
