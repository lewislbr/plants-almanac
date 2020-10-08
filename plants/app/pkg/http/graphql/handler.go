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

// Start initalizes the GraphQL endpoints
func Start() error {
	godotenv.Load()

	router := httprouter.New()
	port := os.Getenv("PLANTS_APP_PORT")
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

	router.POST("/plants", httpWrapper(graphqlHandler))

	if isDevelopment {
		router.GET("/playground", httpWrapper(playgroundHandler))
	}

	fmt.Println("Server ready ✅")

	var err error
	if isDevelopment {
		err = http.ListenAndServe(":"+port, corsWrapper(router))
	} else {
		err = http.ListenAndServeTLS(":443", "etc/tls/server.crt", "etc/tls/server.key", corsWrapper(router))
	}
	if err != nil {
		return err
	}

	return nil
}

func httpWrapper(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}

func corsWrapper(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var origin string
		if isDevelopment {
			origin = "*"
		} else {
			origin = os.Getenv("WEB_PRODUCTION_URL")
		}

		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin")
		w.Header().Add("Access-Control-Max-Age", "86400")
		w.Header().Add("Access-Control-Allow-Origin", origin)

		if !isDevelopment {
			w.Header().Add("Content-Security-Policy", "default-src 'self'")
			w.Header().Add("Strict-Transport-Security",
				"max-age=63072000; includeSubDomains; preload")
		}
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	}
}
