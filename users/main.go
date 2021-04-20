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
	DBURI  string
	Domain string
	Port   string
	Secret string
}

func main() {
	env := getEnvVars()
	str := storage.New()
	db, err := str.Connect(env.DBURI)
	if err != nil {
		log.Panic(err)
	}

	rep := storage.NewRepository(db)
	crs := create.NewService(rep)
	gns := generate.NewService(env.Secret)
	ans := authenticate.NewService(gns, rep)
	azs := authorize.NewService(env.Secret)
	srv := server.New(crs, ans, azs, gns, env.Port, env.Domain)

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
		DBURI:  get("USERS_DATABASE_URI"),
		Domain: get("APP_DOMAIN"),
		Port:   get("USERS_PORT"),
		Secret: get("USERS_SECRET"),
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

	str.Disconnect()

	err := srv.Stop(ctx)
	if err != nil {
		fmt.Println(err)
	}
}
