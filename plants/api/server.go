package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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
		return errors.Wrap(err, "")
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

	fmt.Println("Plants server ready ✅")

	err = Server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return errors.Wrap(err, "")
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

func corsMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin")
		w.Header().Add("Access-Control-Allow-Methods", "POST")
		w.Header().Add("Access-Control-Allow-Origin", os.Getenv("WEB_URL"))
		w.Header().Add("Access-Control-Max-Age", "86400")
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)

			return
		}

		h.ServeHTTP(w, r)
	}
}

func authorizationMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest("GET", os.Getenv("USERS_AUTHORIZATION_URL"), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}

		for _, cookie := range r.Cookies() {
			if cookie.Name == "st" {
				req.Header.Add("Authorization", "Bearer "+cookie.Value)
			}
		}

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
		if res.StatusCode != http.StatusOK {
			w.WriteHeader(res.StatusCode)

			return
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}

		uid = string(body)

		h.ServeHTTP(w, r)
	}
}
