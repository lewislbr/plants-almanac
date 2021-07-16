package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lewislbr/plantdex/users/server"
	"lewislbr/plantdex/users/storage/tokenstore"
	"lewislbr/plantdex/users/storage/userstore"
	"lewislbr/plantdex/users/token"
	"lewislbr/plantdex/users/user"
)

type envVars struct {
	AppDomain    string
	GCPProjectId string
	TokenSecret  string
	WebURL       string
}

func main() {
	if svr, err := run(); err != nil {
		log.Printf("Error starting app: %v\n", err)

		err := cleanUp(svr)
		if err != nil {
			log.Printf("Error cleaning up: %v\n", err)
		}

		os.Exit(1)
	}
}

func run() (*server.Server, error) {
	env := getEnvVars()
	userStore := userstore.New()
	userDB, err := userStore.Connect(env.GCPProjectId)
	if err != nil {
		return nil, fmt.Errorf("error connecting user database: %v\n", err)
	}

	tokenStore := tokenstore.New()
	tokenDB, err := tokenStore.Connect(env.GCPProjectId)
	if err != nil {
		return nil, fmt.Errorf("error connecting token database: %v\n", err)
	}

	userRepo := userstore.NewRepository(userDB)
	tokenRepo := tokenstore.NewRepository(tokenDB)
	userSvc := user.NewService(userRepo)
	tokenSvc := token.NewService(env.TokenSecret, tokenRepo)
	userSvr := server.New(userSvc, tokenSvc, env.AppDomain, env.WebURL)

	go gracefulShutdown(userSvr, userStore)

	err = userSvr.Start()
	if err != nil {
		return userSvr, fmt.Errorf("error starting users server: %v\n", err)
	}

	return nil, nil
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
		AppDomain:    get("APP_DOMAIN"),
		GCPProjectId: get("GCP_PROJECT_ID"),
		TokenSecret:  get("USERS_SECRET"),
		WebURL:       get("WEB_URL"),
	}
}

func gracefulShutdown(svr *server.Server, db *userstore.Driver) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	err := cleanUp(svr)
	if err != nil {
		log.Printf("Error cleaning up: %v\n", err)
	}
}

func cleanUp(svr *server.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := svr.Stop(ctx)
	if err != nil {
		return fmt.Errorf("error stopping server: %v\n", err)
	}

	return nil
}
