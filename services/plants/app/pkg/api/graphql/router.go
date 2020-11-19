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

// Start initalizes the GraphQL API
func Start() error {
	godotenv.Load(".plants-env")

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

	router := httprouter.New()

	router.POST("/plants", responseMiddleware(graphqlHandler))

	if isDevelopment {
		router.GET("/playground", responseMiddleware(playgroundHandler))
	}

	fmt.Println("Plants API ready ✅")

	port := os.Getenv("PLANTS_APP_PORT")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return err
	}

	return nil
}

func responseMiddleware(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}
