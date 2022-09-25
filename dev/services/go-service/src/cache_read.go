package main

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"github.com/rs/zerolog/log"
	"strconv"
)

type readResult struct {
	value int64
	err   error
}

func cacheRead(c chan readResult) readResult {

	result := readResult{}
	if result.err != nil {
		c <- result
		return result
	}

	v, err := cache.Get(context.Background(), getCacheKey()).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			c <- result
		} else {
			log.Err(err).Msgf("cache fail: %s", err.Error())
			result.err = err
		}
	} else {
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Err(err).Msgf("cache fail: %s", err.Error())
			result.err = err
		} else {
			result.value = value
		}
	}
	c <- result
	return result
}
