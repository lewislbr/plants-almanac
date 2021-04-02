package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"plants/add"
	"plants/delete"
	"plants/edit"
	"plants/list"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/pkg/errors"
)

var server = &http.Server{}

func setUpRouter(h *handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	r.Use(cors.Handler(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Origin"},
		AllowedMethods:   []string{"DELETE", "GET", "OPTIONS", "PUT", "POST"},
		AllowedOrigins:   []string{os.Getenv("WEB_URL")},
		MaxAge:           86400,
	}))
	r.Use(headersMiddleware, authorizationMiddleware)

	r.Post("/add", h.Add)
	r.Get("/list", h.ListAll)
	r.Get("/list/{id}", h.ListOne)
	r.Put("/edit/{id}", h.Edit)
	r.Delete("/delete/{id}", h.Delete)

	return r
}

func Start(as add.AddService, ls list.ListService, es edit.EditService, ds delete.DeleteService) error {
	handler := NewHandler(as, ls, es, ds)
	router := setUpRouter(handler)
	port := os.Getenv("PLANTS_PORT")

	server.Addr = ":" + port
	server.Handler = router
	server.IdleTimeout = 120 * time.Second
	server.MaxHeaderBytes = 1 << 20 // 1 MB
	server.ReadTimeout = 5 * time.Second
	server.WriteTimeout = 10 * time.Second

	fmt.Println("Plants server ready ✅")

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
