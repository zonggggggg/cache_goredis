package cache

import (
	"context"

	log "cache_goredis/logger"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	*redis.Client
	ctx context.Context
}

// NewRedis creates a new Redis client.
//
// It takes a context.Context as a parameter.
// It returns a *redis.Client.
func NewRedis(ctx context.Context) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "admin",
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Error("redis client.Ping failed", err)
		panic(err)
	}

	return &Redis{Client: client, ctx: ctx}
}

// SetKeyValue sets a key-value pair in the Redis cache.
//
// Parameters:
// - key: the key to set in the cache.
// - value: the value to associate with the key.
//
// Returns:
// - error: an error if the operation fails.
func (r *Redis) SetKeyValue(key string, value interface{}) error {
	err := r.Set(r.ctx, key, value, 0).Err()
	if err != nil {
		log.Warn("redis client.Set failed", err)
	}
	return err
}

// GetValue returns the value associated with the given key from the Redis cache.
//
// Parameters:
// - key: the key to retrieve the value for.
//
// Returns:
// - interface{}: the value associated with the key.
// - error: an error if the retrieval fails.
func (r *Redis) GetValue(key string) (interface{}, error) {
	return r.Get(r.ctx, key).Result()
}

func (r *Redis) DeleteKey(key string) error {
	err := r.Del(r.ctx, key).Err()
	if err != nil {
		log.Warn("redis client.Del failed", err)
	}
	return err
}
