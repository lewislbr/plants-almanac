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

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

var a = add.NewService(&storage.Inmem{})
var l = list.NewService(&storage.Inmem{})
var e = edit.NewService(&storage.Inmem{})
var d = delete.NewService(&storage.Inmem{})

// Start initializes the REST API
func Start() error {
	godotenv.Load(".users-env")

	router := httprouter.New()
	port := os.Getenv("USERS_APP_PORT")
	endpointRoot := "/users"

	router.HandlerFunc("POST", endpointRoot, addUser(a))
	router.HandlerFunc("GET", endpointRoot+"/{id}", listUser(l))
	router.HandlerFunc("PATCH", endpointRoot+"/{id}", editUser(e))
	router.HandlerFunc("DELETE", endpointRoot+"/{id}", deleteUser(d))

	fmt.Println("Users API ready ✅")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return err
	}

	return nil
}
