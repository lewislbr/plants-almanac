package graphql

import (
	"fmt"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

var isDevelopment = os.Getenv("MODE") == "development"

// Start initalizes the GraphQL API
func Start() error {
	godotenv.Load()

	graphqlHandler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     false,
		Playground: false,
	})
	playgroundHandler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		Playground: true,
	})

	router := httprouter.New()

	router.POST("/plants", responseMiddleware(graphqlHandler))

	if isDevelopment {
		router.GET("/playground", responseMiddleware(playgroundHandler))
	}

	fmt.Println("Plants API ready ✅")

	var err error
	if isDevelopment {
		port := os.Getenv("PLANTS_APP_PORT")

		err = http.ListenAndServe(":"+port, corsMiddleware(router))
	} else {
		err = http.ListenAndServeTLS(
			":443",
			"etc/tls/server.crt",
			"etc/tls/server.key",
			corsMiddleware(router),
		)
	}
	if err != nil {
		return err
	}

	return nil
}

func responseMiddleware(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}

func corsMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var origin string
		if isDevelopment {
			origin = "*"
		} else {
			origin = os.Getenv("WEB_PRODUCTION_URL")
		}

		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin")
		w.Header().Add("Access-Control-Allow-Origin", origin)
		w.Header().Add("Access-Control-Max-Age", "86400")

		if !isDevelopment {
			w.Header().Add("Content-Security-Policy", "default-src 'self'")
			w.Header().Add(
				"Strict-Transport-Security",
				"max-age=63072000; includeSubDomains; preload",
			)
		}
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	}
}
