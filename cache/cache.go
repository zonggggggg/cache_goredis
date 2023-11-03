package cache

import "context"

const (
	REDIS = "redis"
)

type ICache interface {
	GetValue(key string) (interface{}, error)
	SetKeyValue(key string, value interface{}) error
	DeleteKey(key string) error
}

// NewCache creates a new cache instance based on the specified option.
//
// ctx: The context.Context object for the cache.
// option: The option string indicating the type of cache to create.
// Returns an ICache interface.
func NewCache(ctx context.Context, option string) ICache {
	switch option {
	case REDIS:
		return NewRedis(ctx)
	}
	return nil
}
