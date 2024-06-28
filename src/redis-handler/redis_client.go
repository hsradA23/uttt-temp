package redis_handler

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background() // This is a default context

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
