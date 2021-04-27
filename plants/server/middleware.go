package server

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
)

type uid string

const contextId uid = "uid"

func authorizationMiddleware(url string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var token string

			for _, cookie := range r.Cookies() {
				if cookie.Name == "st" {
					token = cookie.Value
				}
			}

			res, err := http.Post(url+"/authorization", "text/plain", bytes.NewBuffer([]byte(token)))
			if err != nil {
				http.Error(w, "something went wrong", http.StatusInternalServerError)

				log.Printf("%+v\n", err)

				return
			}
			if res.StatusCode != http.StatusOK {
				body, err := io.ReadAll(res.Body)
				if err != nil {
					log.Printf("%+v\n", err)
				}

				w.WriteHeader(res.StatusCode)
				_, err = w.Write(body)
				if err != nil {
					log.Printf("%+v\n", err)
				}

				return
			}

			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				http.Error(w, "something went wrong", http.StatusInternalServerError)

				log.Printf("%+v\n", err)

				return
			}

			uid := string(body)
			ctx := context.WithValue(r.Context(), contextId, uid)

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
