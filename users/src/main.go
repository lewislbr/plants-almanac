package main

import (
	"log"

	"users/src/api"
	"users/src/authenticate"
	"users/src/authorize"
	"users/src/create"
	"users/src/generate"
	"users/src/storage"
)

func main() {
	db := storage.ConnectDatabase()
	repository := storage.NewRepository(db)
	createService := create.NewCreateService(repository)
	generateService := generate.NewGenerateService()
	authenticateService := authenticate.NewAuthenticateService(generateService, repository)
	authorizeService := authorize.NewAuthorizeService()

	if err := api.Start(createService, authenticateService, authorizeService, generateService); err != nil {
		log.Fatalln(err)
	}
}
