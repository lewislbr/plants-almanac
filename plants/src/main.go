package main

import (
	"fmt"
	"net/http"

	"plants/src/graphql/schema"

	"github.com/graphql-go/handler"
)

func main() {
	handler := handler.New(&handler.Config{
		Schema:     &schema.Schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", corsHandler(handler))

	port := ":4040"

	fmt.Printf("Server ready at http://localhost%v ✅\n", port)

	http.ListenAndServe(port, nil)
}

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin")
		w.Header().Add("Access-Control-Max-Age", "86400")
		w.Header().Add("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
		} else {
			h.ServeHTTP(w, r)
		}
	}
}
