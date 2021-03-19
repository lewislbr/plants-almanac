package server

import (
	"io"
	"log"
	"net/http"
	"os"
)

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
