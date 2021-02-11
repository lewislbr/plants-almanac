package main

import (
	"log"

	"users/src/api"
	"users/src/authenticate"
	"users/src/authorize"
	"users/src/create"
	"users/src/storage"
)

func main() {
	db := storage.ConnectDatabase()
	repository := storage.NewRepository(db)
	createService := create.NewCreateService(repository)
	authenticateService := authenticate.NewAuthenticateService(repository)
	authorizeService := authorize.NewAuthorizeService()

	if err := api.Start(createService, authenticateService, authorizeService); err != nil {
		log.Fatalln(err)
	}
}
