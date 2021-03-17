package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"users/authenticate"
	"users/authorize"
	"users/create"
	"users/generate"
	"users/server"
	"users/storage"
)

var db, err = storage.ConnectDatabase()

func main() {
	defer func() {
		r := recover()
		if r != nil {
			cleanUp()
			debug.PrintStack()
			os.Exit(1)
		}
	}()

	if err != nil {
		log.Panic(err)
	}

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

func cleanUp() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if db != nil {
		err := storage.DisconnectDatabase(ctx, db)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = server.Stop(ctx)
	if err != nil {
		fmt.Println(err)
	}
}

func gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	cleanUp()
}
