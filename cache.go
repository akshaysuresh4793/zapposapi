package main

import (
	"github.com/go-redis/redis"
	"os"
	"fmt"
)

var cache *redis.Client
func initialize() {
	cacheHost := "localhost"
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
	err := cache.Set(key, value, 0).Err()
	handleError(err)
}

func get(key string) string {
	value, err := client.Get(key).Result()
	if err == redis.Nil {
		// do nothing
		return ""
	}
	handleError(err)
	if err != nil {
		return ""
	}
	return value
}
