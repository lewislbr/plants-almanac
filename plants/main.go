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

	"plants/add"
	"plants/delete"
	"plants/edit"
	"plants/list"
	"plants/server"
	"plants/storage"
)

type envVars struct {
	AuthURL  string
	Database string
	MongoURI string
	Port     string
	WebURL   string
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
	db, err := storage.Connect(env.MongoURI, env.Database)
	if err != nil {
		log.Panic(err)
	}

	r := storage.NewRepository(db)
	as := add.NewService(r)
	ls := list.NewService(r)
	es := edit.NewService(ls, r)
	ds := delete.NewService(r)

	server.New(as, ls, es, ds, env.Port, env.AuthURL, env.WebURL)

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
		AuthURL:  get("USERS_URL"),
		Database: get("PLANTS_DATABASE_NAME"),
		MongoURI: get("PLANTS_MONGODB_URI"),
		Port:     get("PLANTS_PORT"),
		WebURL:   get("WEB_URL"),
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
