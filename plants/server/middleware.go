package server

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

type userID string

const contextId userID = "userID"

func headersMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		h.ServeHTTP(w, r)
	})
}

func authzMiddleware(authUrl string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var token string

			for _, cookie := range r.Cookies() {
				if cookie.Name == "st" {
					token = cookie.Value
				}
			}

			client := &http.Client{
				Timeout: time.Second * 10,
			}
			res, err := client.Post(authUrl+"/authorization", "text/plain", bytes.NewBuffer([]byte(token)))
			if err != nil {
				http.Error(w, "Something went wrong", http.StatusInternalServerError)

				log.Printf("Error requesting authorization: %v\n", err)

				return
			}

			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				http.Error(w, "Something went wrong", http.StatusInternalServerError)

				log.Printf("Error reading authorization response: %v\n", err)

				return
			}

			if res.StatusCode != http.StatusOK {
				w.WriteHeader(res.StatusCode)
				_, err = w.Write(body)
				if err != nil {
					log.Printf("Error forwarding authorization response: %v\n", err)
				}

				return
			}

			userID := string(body)
			ctx := context.WithValue(r.Context(), contextId, userID)

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
