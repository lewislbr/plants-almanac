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
	env := getEnvVars()
	str := storage.New()
	db, err := str.Connect(env.MongoURI, env.Database, env.Collection)
	if err != nil {
		log.Panic(err)
	}

	rep := storage.NewRepository(db)
	crs := create.NewService(rep)
	gns := generate.NewService(env.Secret)
	ans := authenticate.NewService(gns, rep)
	azs := authorize.NewService(env.Secret)
	srv := server.New(crs, ans, azs, gns, env.Port, env.WebURL)

	go gracefulShutdown(srv, str)

	err = srv.Start()
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		r := recover()
		if r != nil {
			cleanUp(srv, str)
			debug.PrintStack()
			os.Exit(1)
		}
	}()
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

func gracefulShutdown(srv *server.Server, str *storage.Storage) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	cleanUp(srv, str)
}

func cleanUp(srv *server.Server, str *storage.Storage) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := str.Disconnect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	err = srv.Stop(ctx)
	if err != nil {
		fmt.Println(err)
	}
}
