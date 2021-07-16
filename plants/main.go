package main

import (
	"context"
	"fmt"
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
	if svr, db, err := start(); err != nil {
		log.Printf("Error starting app: %v\n", err)

		err := cleanUp(svr, db)
		if err != nil {
			log.Printf("Error cleaning up: %v\n", err)
		}

		os.Exit(1)
	}
}

func start() (*server.Server, *plantstore.Driver, error) {
	env := getEnvVars()
	plantStore := plantstore.New()
	plantDB, err := plantStore.Connect(env.MongoURI, env.Database)
	if err != nil {
		return nil, plantStore, fmt.Errorf("error connecting plants database: %v\n", err)
	}

	plantRepo := plantstore.NewRepository(plantDB)
	plantSvc := plant.NewService(plantRepo)
	plantSvr := server.New(plantSvc, env.AuthURL, env.WebURL)

	go gracefulShutdown(plantSvr, plantStore)

	err = plantSvr.Start()
	if err != nil {
		return plantSvr, plantStore, fmt.Errorf("error starting plants server: %v\n", err)
	}

	return nil, nil, nil
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

	err := cleanUp(svr, db)
	if err != nil {
		log.Printf("Error cleaning up: %v\n", err)
	}
}

func cleanUp(svr *server.Server, db *plantstore.Driver) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := db.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("error disconnecting plants database: %v\n", err)
	}

	err = svr.Stop(ctx)
	if err != nil {
		return fmt.Errorf("error disconnecting plants server: %v\n", err)
	}

	return nil
}
