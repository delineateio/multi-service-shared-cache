package main

import (
	"context"
)

func cacheWrite(increment int64) {
	cache.IncrBy(context.Background(), getCacheKey(), increment)
}
