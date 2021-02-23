package main

import (
	"log"

	"users/api"
	"users/authenticate"
	"users/authorize"
	"users/create"
	"users/generate"
	"users/storage"
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
