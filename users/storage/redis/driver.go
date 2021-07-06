package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type Driver struct {
	client *redis.Client
}

func New() *Driver {
	return &Driver{}
}

func (d *Driver) Connect(url, pass string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		DB:       0,
		Password: pass,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("error pinging Redis client: %w", err)
	}

	log.Println("Users cache ready ✅")

	d.client = client

	return d.client, nil
}

func (d *Driver) Disconnect() error {
	log.Println("Disconnecting cache...")

	err := d.client.Close()

	return fmt.Errorf("error disconnecting Redis driver: %w", err)
}
