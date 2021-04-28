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
	"users/info"
	"users/revoke"
	"users/server"
	"users/storage/postgres"
	"users/storage/redis"
)

type envVars struct {
	PostgresURI string
	AppDomain   string
	RedisPass   string
	RedisURL    string
	TokenSecret string
}

func main() {
	env := getEnvVars()
	postgresDriver := postgres.New()
	postgresDB, err := postgresDriver.Connect(env.PostgresURI)
	if err != nil {
		log.Panic(err)
	}

	redisDriver := redis.New()
	redisCache, err := redisDriver.Connect(env.RedisURL, env.RedisPass)
	if err != nil {
		log.Panic(err)
	}

	postgresRepo := postgres.NewRepository(postgresDB)
	redisRepo := redis.NewRepository(redisCache)
	createSvc := create.NewService(postgresRepo)
	generateSvc := generate.NewService(env.TokenSecret)
	authenticateSvc := authenticate.NewService(generateSvc, postgresRepo)
	authorizeSvc := authorize.NewService(env.TokenSecret, redisRepo)
	revokeSvc := revoke.NewService(env.TokenSecret, redisRepo)
	infoSvc := info.NewService(postgresRepo)
	httpServer := server.New(createSvc, authenticateSvc, authorizeSvc, generateSvc, revokeSvc, infoSvc, env.AppDomain)

	go gracefulShutdown(httpServer, postgresDriver, redisDriver)

	err = httpServer.Start()
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		r := recover()
		if r != nil {
			cleanUp(httpServer, postgresDriver, redisDriver)
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
		PostgresURI: get("USERS_DATABASE_URI"),
		AppDomain:   get("APP_DOMAIN"),
		RedisPass:   get("USERS_REDIS_PASSWORD"),
		RedisURL:    get("USERS_REDIS_URL"),
		TokenSecret: get("USERS_SECRET"),
	}
}

func gracefulShutdown(http *server.Server, postgres *postgres.Driver, redis *redis.Driver) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	cleanUp(http, postgres, redis)
}

func cleanUp(http *server.Server, postgres *postgres.Driver, redis *redis.Driver) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	postgres.Disconnect()

	err := redis.Disconnect()
	if err != nil {
		fmt.Println(err)
	}

	err = http.Stop(ctx)
	if err != nil {
		fmt.Println(err)
	}
}
