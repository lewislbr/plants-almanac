package server

import (
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
			req, err := http.NewRequest("GET", url+"/authorization", nil)
			if err != nil {
				http.Error(w, "something went wrong", http.StatusInternalServerError)

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
				http.Error(w, "something went wrong", http.StatusInternalServerError)

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
