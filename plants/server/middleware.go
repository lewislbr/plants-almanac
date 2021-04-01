package server

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

type uid string

const contextId uid = "uid"

func headersMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		h.ServeHTTP(w, r)
	})
}

func authorizationMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		uid := string(body)
		ctx := context.WithValue(r.Context(), contextId, uid)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
