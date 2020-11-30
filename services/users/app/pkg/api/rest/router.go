package rest

import (
	"fmt"
	"net/http"
	"os"

	"users/pkg/add"
	"users/pkg/delete"
	"users/pkg/edit"
	"users/pkg/list"
	"users/pkg/storage/inmem"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

var a = add.NewService(&inmem.Storage{})
var l = list.NewService(&inmem.Storage{})
var e = edit.NewService(&inmem.Storage{})
var d = delete.NewService(&inmem.Storage{})

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
