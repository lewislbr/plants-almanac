package main

import (
	"context"
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
	env := getEnvVars()
	userStore := userstore.New()
	userDB, err := userStore.Connect(env.PostgresURI)
	if err != nil {
		log.Panicf("Error connecting user database: %v\n", err)
	}

	tokenStore := tokenstore.New()
	tokenDB, err := tokenStore.Connect(env.GCPProjectId)
	if err != nil {
		log.Panicf("Error connecting token database: %v\n", err)
	}

	userRepo := userstore.NewRepository(userDB)
	tokenRepo := tokenstore.NewRepository(tokenDB)
	userSvc := user.NewService(userRepo)
	tokenSvc := token.NewService(env.TokenSecret, tokenRepo)
	userSvr := server.New(userSvc, tokenSvc, env.AppDomain, env.WebURL)

	go gracefulShutdown(userSvr, userStore)

	err = userSvr.Start()
	if err != nil {
		log.Panicf("Error starting server: %v\n", err)
	}

	defer func() {
		r := recover()
		if r != nil {
			cleanUp(userSvr, userStore)

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

	cleanUp(svr, db)
}

func cleanUp(svr *server.Server, db *userstore.Driver) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := svr.Stop(ctx)
	if err != nil {
		log.Printf("Error stopping server: %v\n", err)
	}
}
