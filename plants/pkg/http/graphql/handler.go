package graphql

import (
	"fmt"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

// Start initalizes the GraphQL endpoints
func Start() error {
	godotenv.Load()

	router := httprouter.New()
	port := os.Getenv("PORT")
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

	router.POST("/api", httpWrapper(graphqlHandler))
	router.GET("/api/playground", httpWrapper(playgroundHandler))

	fmt.Printf("Server ready at http://localhost:%v ✅\n", port)

	err := http.ListenAndServe(":"+port, corsWrapper(router))
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
