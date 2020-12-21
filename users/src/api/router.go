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

// Start initializes the REST API
func Start() error {
	router := httprouter.New()
	port := os.Getenv("USERS_PORT")

	router.HandlerFunc("POST", "/", addUser(addService))
	router.HandlerFunc("GET", "/:id", listUser(listService))
	router.HandlerFunc("PATCH", "/:id", editUser(editService))
	router.HandlerFunc("DELETE", "/:id", deleteUser(deleteService))

	fmt.Println("Users API ready ✅")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return err
	}

	return nil
}
