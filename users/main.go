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

type envVars struct {
	Collection string
	Database   string
	MongoURI   string
	Port       string
	Secret     string
	WebURL     string
}

var (
	env     = getEnvVars()
	db, err = storage.ConnectDatabase(env.MongoURI, env.Database, env.Collection)
)

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
	gs := generate.NewGenerateService(env.Secret)
	ns := authenticate.NewAuthenticateService(gs, r)
	zs := authorize.NewAuthorizeService(env.Secret)

	go gracefulShutdown()

	err := server.Start(cs, ns, zs, gs, env.Port, env.WebURL)
	if err != nil {
		log.Panic(err)
	}
}

func getEnvVars() *envVars {
	get := func(k string) string {
		v, set := os.LookupEnv(k)
		if !set || v == "" {
			log.Fatalf("%q environment variable not set.\n", k)
		}

		return v
	}

	return &envVars{
		Collection: get("USERS_COLLECTION_NAME"),
		Database:   get("USERS_DATABASE_NAME"),
		MongoURI:   get("USERS_MONGODB_URI"),
		Port:       get("USERS_PORT"),
		Secret:     get("USERS_JWT_SECRET"),
		WebURL:     get("WEB_URL"),
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
