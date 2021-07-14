package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lewislbr/plantdex/plants/plant"
	"lewislbr/plantdex/plants/server"
	"lewislbr/plantdex/plants/storage/plantstore"
)

type envVars struct {
	AuthURL  string
	Database string
	MongoURI string
	WebURL   string
}

func main() {
	env := getEnvVars()
	plantStore := plantstore.New()
	plantDB, err := plantStore.Connect(env.MongoURI, env.Database)
	if err != nil {
		log.Panicf("Error connecting plants database: %v\n", err)
	}

	plantRepo := plantstore.NewRepository(plantDB)
	plantSvc := plant.NewService(plantRepo)
	plantSvr := server.New(plantSvc, env.AuthURL, env.WebURL)

	go gracefulShutdown(plantSvr, plantStore)

	err = plantSvr.Start()
	if err != nil {
		log.Panicf("Error starting server: %v\n", err)
	}

	defer func() {
		r := recover()
		if r != nil {
			cleanUp(plantSvr, plantStore)

			os.Exit(1)
		}
	}()
}

func getEnvVars() *envVars {
	get := func(k string) string {
		v, set := os.LookupEnv(k)
		if !set || v == "" {
			log.Fatalf("%q environment variable not set\n", k)
		}

		return v
	}

	return &envVars{
		AuthURL:  get("USERS_URL"),
		Database: get("PLANTS_DATABASE_NAME"),
		MongoURI: get("PLANTS_DATABASE_URI"),
		WebURL:   get("WEB_URL"),
	}
}

func gracefulShutdown(svr *server.Server, db *plantstore.Driver) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	cleanUp(svr, db)
}

func cleanUp(svr *server.Server, db *plantstore.Driver) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := db.Disconnect(ctx)
	if err != nil {
		log.Printf("Error disconnecting plants database: %v\n", err)
	}

	err = svr.Stop(ctx)
	if err != nil {
		log.Printf("Error stopping server: %v\n", err)
	}
}
