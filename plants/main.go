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
	DBURI    string
	Port     string
	WebURL   string
}

func main() {
	env := getEnvVars()
	str := storage.New()
	db, err := str.Connect(env.DBURI, env.Database)
	if err != nil {
		log.Panic(err)
	}

	rep := storage.NewRepository(db)
	ads := add.NewService(rep)
	lss := list.NewService(rep)
	eds := edit.NewService(lss, rep)
	dls := delete.NewService(rep)
	srv := server.New(ads, lss, eds, dls, env.Port, env.AuthURL, env.WebURL)

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
		AuthURL:  get("USERS_URL"),
		Database: get("PLANTS_DATABASE_NAME"),
		DBURI:    get("PLANTS_DATABASE_URI"),
		Port:     get("PLANTS_PORT"),
		WebURL:   get("WEB_URL"),
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
