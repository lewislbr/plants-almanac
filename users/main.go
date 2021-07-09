package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lewislbr/plantdex/users/server"
	"lewislbr/plantdex/users/storage/postgres"
	"lewislbr/plantdex/users/storage/redis"
	"lewislbr/plantdex/users/token"
	"lewislbr/plantdex/users/user"
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
		log.Panicf("Error connecting database: %v\n", err)
	}

	redisDriver := redis.New()
	redisCache, err := redisDriver.Connect(env.RedisURL, env.RedisPass)
	if err != nil {
		log.Panicf("Error connecting cache: %v\n", err)
	}

	postgresRepo := postgres.NewRepository(postgresDB)
	redisRepo := redis.NewRepository(redisCache)
	userSvc := user.NewService(postgresRepo)
	tokenSvc := token.NewService(env.TokenSecret, redisRepo)
	userSvr := server.New(userSvc, tokenSvc, env.AppDomain)

	go gracefulShutdown(userSvr, postgresDriver, redisDriver)

	err = userSvr.Start()
	if err != nil {
		log.Panicf("Error starting server: %v\n", err)
	}

	defer func() {
		r := recover()
		if r != nil {
			cleanUp(userSvr, postgresDriver, redisDriver)

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
		PostgresURI: get("USERS_DATABASE_URI"),
		AppDomain:   get("APP_DOMAIN"),
		RedisPass:   get("USERS_REDIS_PASSWORD"),
		RedisURL:    get("USERS_REDIS_URL"),
		TokenSecret: get("USERS_SECRET"),
	}
}

func gracefulShutdown(svr *server.Server, db *postgres.Driver, cache *redis.Driver) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	cleanUp(svr, db, cache)
}

func cleanUp(svr *server.Server, db *postgres.Driver, cache *redis.Driver) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db.Disconnect()

	err := cache.Disconnect()
	if err != nil {
		log.Printf("Error disconnecting cache: %v\n", err)
	}

	err = svr.Stop(ctx)
	if err != nil {
		log.Printf("Error stopping server: %v\n", err)
	}
}
