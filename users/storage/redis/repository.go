package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type repository struct {
	cache *redis.Client
}

func NewRepository(cache *redis.Client) *repository {
	return &repository{cache}
}

func (r *repository) Add(tokenId string) error {
	err := r.cache.Set(context.Background(), tokenId, 0, 7*24*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("error adding user: %w", err)
	}

	return nil
}

func (r *repository) CheckExists(tokenId string) error {
	err := r.cache.Get(context.Background(), tokenId).Err()
	if err != nil {
		return fmt.Errorf("error checking user: %w", err)
	}

	return nil
}
