package api

import (
	"fmt"
	"net/http"
	"os"

	"users/src/add"
	"users/src/delete"
	"users/src/edit"
	"users/src/list"
	"users/src/storage"

	"github.com/julienschmidt/httprouter"
)

var addService = add.NewService(&storage.Inmem{})
var listService = list.NewService(&storage.Inmem{})
var editService = edit.NewService(&storage.Inmem{})
var deleteService = delete.NewService(&storage.Inmem{})
var isDevelopment = os.Getenv("MODE") == "development"

// Start initializes the REST API
func Start() error {
	router := httprouter.New()
	port := os.Getenv("USERS_PORT")

	router.HandlerFunc("POST", "/", addUser(addService))
	router.HandlerFunc("GET", "/:id", listUser(listService))
	router.HandlerFunc("PATCH", "/:id", editUser(editService))
	router.HandlerFunc("DELETE", "/:id", deleteUser(deleteService))

	fmt.Println("Users API ready ✅")

	err := http.ListenAndServe(":"+port, corsMiddleware(router))
	if err != nil {
		return err
	}

	return nil
}

func corsMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var origin string
		if isDevelopment {
			origin = "*"
		} else {
			origin = os.Getenv("WEB_PRODUCTION_URL")
		}

		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin")
		w.Header().Add(
			"Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE",
		)
		w.Header().Add("Access-Control-Allow-Origin", origin)
		w.Header().Add("Access-Control-Max-Age", "86400")
		w.Header().Add(
			"Strict-Transport-Security",
			"max-age=63072000; includeSubDomains; preload",
		)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)

			return
		}

		h.ServeHTTP(w, r)
	}
}
