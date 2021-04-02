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

func main() {
	defer func() {
		r := recover()
		if r != nil {
			cleanUp()
			debug.PrintStack()
			os.Exit(1)
		}
	}()

	env := getEnvVars()
	db, err := storage.Connect(env.MongoURI, env.Database, env.Collection)
	if err != nil {
		log.Panic(err)
	}

	r := storage.NewRepository(db)
	cs := create.NewService(r)
	gs := generate.NewService(env.Secret)
	ns := authenticate.NewService(gs, r)
	zs := authorize.NewService(env.Secret)

	server.New(cs, ns, zs, gs, env.Port, env.WebURL)

	go gracefulShutdown()

	err = server.Start()
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

func gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	cleanUp()
}

func cleanUp() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := storage.Disconnect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	err = server.Stop(ctx)
	if err != nil {
		fmt.Println(err)
	}
}
