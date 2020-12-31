package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/julienschmidt/httprouter"
)

var isDevelopment = os.Getenv("MODE") == "development"
var uid string

// Start initalizes the GraphQL API
func Start() error {
	graphqlHandler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     false,
		Playground: false,
		RootObjectFn: func(ctx context.Context, r *http.Request) map[string]interface{} {
			return map[string]interface{}{
				"uid": uid,
			}
		},
	})
	playgroundHandler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		Playground: true,
	})

	router := httprouter.New()

	router.Handler("POST", "/", graphqlHandler)

	if isDevelopment {
		router.Handler("GET", "/playground", playgroundHandler)
	}

	fmt.Println("Plants API ready ✅")

	port := os.Getenv("PLANTS_PORT")
	err := http.ListenAndServe(":"+port, corsMiddleware(authorizationMiddleware(router)))
	if err != nil {
		return err
	}

	return nil
}

func corsMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var origin string
		if isDevelopment {
			origin = os.Getenv("WEB_DEVELOPMENT_URL")
		} else {
			origin = os.Getenv("WEB_PRODUCTION_URL")
		}

		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin")
		w.Header().Add("Access-Control-Allow-Methods", "POST")
		w.Header().Add("Access-Control-Allow-Origin", origin)
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
		var authURL string
		if isDevelopment {
			authURL = os.Getenv("USERS_AUTHORIZATION_DEVELOPMENT_URL")
		} else {
			authURL = os.Getenv("USERS_AUTHORIZATION_PRODUCTION_URL")
		}

		req, err := http.NewRequest("GET", authURL, nil)
		if err != nil {
			fmt.Println(err)
		}

		for _, cookie := range r.Cookies() {
			if cookie.Name == "st" {
				req.Header.Add("Authorization", "Bearer "+cookie.Value)
			}
		}

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		if res.StatusCode != http.StatusOK {
			w.WriteHeader(res.StatusCode)

			return
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		uid = string(body)

		h.ServeHTTP(w, r)
	}
}
