package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/rs/zerolog/log"
	"os"
)

var cache *redis.Client

func getCacheHostname(def string) string {

	value := os.Getenv("CACHE_HOSTNAME")
	if value == "" {
		value = def
	}
	return value
}

func getConnectionString() string {

	host := getCacheHostname("localhost")
	port := getString("cache.port", "6379")
	db := getString("cache.db", "0")
	return fmt.Sprintf("redis://:@%s:%s/%s", host, port, db)
}

func getCacheKey() string {
	return getString("cache.key", "count")
}

func connectToCache() error {
	opt, err := redis.ParseURL(getConnectionString())
	if err != nil {
		return err
	}
	cache = redis.NewClient(opt)
	return pingCache()
}

func pingCache() error {
	_, err := cache.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	log.Info().Msgf("cache connection success: %s", cache.Options().Addr)
	return nil
}
