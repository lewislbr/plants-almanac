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
	r := storage.NewRepository(db)
	cs := create.NewCreateService(r)
	gs := generate.NewGenerateService()
	ns := authenticate.NewAuthenticateService(gs, r)
	zs := authorize.NewAuthorizeService()

	if err := api.Start(cs, ns, zs, gs); err != nil {
		log.Fatalln(err)
	}
}
