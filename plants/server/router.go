package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	p "plants/plant"

	"github.com/graphql-go/handler"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

var (
	uid string
	// Server defines the server struct.
	Server = &http.Server{}
)

// Start initalizes the server.
func Start(as p.AddService, ls p.ListService, es p.EditService, ds p.DeleteService) error {
	port := os.Getenv("PLANTS_PORT")
	Server.Addr = ":" + port

	resolver := NewResolver(as, ls, es, ds)
	schema, err := NewSchema(resolver)
	if err != nil {
		return err
	}

	router := httprouter.New()
	handler := handler.New(&handler.Config{
		Schema:     schema,
		Pretty:     false,
		Playground: false,
		RootObjectFn: func(ctx context.Context, r *http.Request) map[string]interface{} {
			return map[string]interface{}{
				"uid": uid,
			}
		},
	})

	router.Handler("POST", "/", handler)

	Server.Handler = corsMiddleware(authorizationMiddleware(router))

	Server.IdleTimeout = 120 * time.Second
	Server.ReadTimeout = 5 * time.Second
	Server.WriteTimeout = 10 * time.Second

	fmt.Println("Plants server ready ✅")

	err = Server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

// Stop stops the server.
func Stop(ctx context.Context) error {
	fmt.Println("Stopping server...")

	err := Server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
