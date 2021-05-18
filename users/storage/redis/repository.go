package redis

import (
	"context"
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
	return r.cache.Set(context.Background(), tokenId, 0, 7*24*time.Hour).Err()
}

func (r *repository) CheckExists(tokenId string) error {
	return r.cache.Get(context.Background(), tokenId).Err()
}
