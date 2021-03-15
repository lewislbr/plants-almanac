package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"users/authenticate"
	"users/authorize"
	"users/create"
	"users/generate"
	"users/server"
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

	err := server.Start(cs, ns, zs, gs)
	if err != nil {
		log.Panic(err)
	}
}

func gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := storage.DisconnectDatabase(ctx, db)
	if err != nil {
		fmt.Print(err)
	}

	err = server.Stop(ctx)
	if err != nil {
		fmt.Print(err)
	}
}
