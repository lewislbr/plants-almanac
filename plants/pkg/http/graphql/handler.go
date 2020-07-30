package graphql

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/handler"
)

// Start initalizes the endpoint where GraphQL is available
func Start() error {
	handler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	port := "4040"

	fmt.Printf("Server ready at http://localhost:%v ✅\n", port)

	err := http.ListenAndServe(":"+port, corsHandler(handler))
	if err != nil {
		return err
	}

	return nil
}

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin")
		w.Header().Add("Access-Control-Max-Age", "86400")
		w.Header().Add("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	}
}
