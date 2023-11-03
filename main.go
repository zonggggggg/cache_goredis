package main

import (
	"cache_goredis/cache"
	log "cache_goredis/logger"
	"context"
)

// main is the entry point of the program.
//
// It initializes a cache using the cache.NewCache function and sets a key-value pair in the cache.
// It then retrieves the value associated with the key from the cache and logs it.
// Finally, it deletes the key from the cache.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cache := cache.NewCache(ctx, cache.REDIS)
	if cache == nil {
		log.Error("cache is nil")
		panic("cache is nil")
	}

	err := cache.SetKeyValue("test key", "test value")
	if err != nil {
		log.Warn("error: ", err)
	}

	val, err := cache.GetValue("test key")
	if err != nil {
		log.Warn("error: ", err)
	} else {
		log.Info("value: ", val)
	}

	err = cache.DeleteKey("test key")
	if err != nil {
		log.Warn("error: ", err)
	}

}
