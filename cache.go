package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"time"
)

var cache *redis.Client

func cacheInitialize() {
	cacheHost := "127.0.0.1"
	cachePort := "6379"
	if len(os.Getenv("CACHEHOST")) > 0 {
		cacheHost = os.Getenv("CACHEHOST")
	}
	if len(os.Getenv("CACHEPORT")) > 0 {
		cachePort = os.Getenv("CACHEPORT")
	}
	fmt.Println("Connecting to Redis - ", cacheHost, ":", cachePort)
	cache = redis.NewClient(&redis.Options{
		Addr:     cacheHost + ":" + cachePort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func set(key string, value string) {
	err := cache.Set(key, value, time.Hour).Err()
	handleError(err)
}

func get(key string) string {
	value, err := cache.Get(key).Result()
	if err == redis.Nil {
		// do nothing
		return ""
	}
	handleError(err)
	if err != nil {
		return ""
	}
	fmt.Println("cache hit")
	return value
}

func delete(key string) {
	_ = cache.Del(key)
}
