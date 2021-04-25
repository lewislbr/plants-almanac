package redis

import (
	"fmt"

	"github.com/go-redis/redis"
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

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	fmt.Println("Users cache ready ✅")

	d.client = client

	return d.client, nil
}

func (d *Driver) Disconnect() error {
	fmt.Println("Disconnecting cache...")

	return d.client.Close()
}
