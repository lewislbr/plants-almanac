package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"users/api"
	"users/authenticate"
	"users/authorize"
	"users/create"
	"users/generate"
	"users/storage"
)

var db = storage.ConnectDatabase()

func main() {
	r := storage.NewRepository(db)
	cs := create.NewCreateService(r)
	gs := generate.NewGenerateService()
	ns := authenticate.NewAuthenticateService(gs, r)
	zs := authorize.NewAuthorizeService()

	go gracefulShutdown()

	err := api.Start(cs, ns, zs, gs)
	if err != nil {
		log.Fatalln(err)
	}
}

func gracefulShutdown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := storage.DisconnectDatabase(ctx, db)
	if err != nil {
		fmt.Print(err)
	}

	err = api.Stop(ctx)
	if err != nil {
		fmt.Print(err)
	}
}
