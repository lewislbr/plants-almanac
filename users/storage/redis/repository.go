package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type repository struct {
	cache *redis.Client
}

func NewRepository(cache *redis.Client) *repository {
	return &repository{cache}
}

func (r *repository) Add(tokenId string) error {
	return r.cache.Set(tokenId, 0, 7*24*time.Hour).Err()
}

func (r *repository) CheckExists(tokenId string) error {
	return r.cache.Get(tokenId).Err()
}
