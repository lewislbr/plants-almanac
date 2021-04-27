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
	"plants/storage/mongo"
)

type envVars struct {
	AuthURL  string
	Database string
	MongoURI string
}

func main() {
	env := getEnvVars()
	mongoDriver := mongo.New()
	mongoDB, err := mongoDriver.Connect(env.MongoURI, env.Database)
	if err != nil {
		log.Panic(err)
	}

	mongoRepo := mongo.NewRepository(mongoDB)
	addSvc := add.NewService(mongoRepo)
	listSvc := list.NewService(mongoRepo)
	editSvc := edit.NewService(listSvc, mongoRepo)
	deleteSvc := delete.NewService(mongoRepo)
	httpServer := server.New(addSvc, listSvc, editSvc, deleteSvc, env.AuthURL)

	go gracefulShutdown(httpServer, mongoDriver)

	err = httpServer.Start()
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		r := recover()
		if r != nil {
			cleanUp(httpServer, mongoDriver)
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
		MongoURI: get("PLANTS_DATABASE_URI"),
	}
}

func gracefulShutdown(http *server.Server, mongo *mongo.Driver) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	cleanUp(http, mongo)
}

func cleanUp(http *server.Server, mongo *mongo.Driver) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := mongo.Disconnect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	err = http.Stop(ctx)
	if err != nil {
		fmt.Println(err)
	}
}
