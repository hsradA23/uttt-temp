package redis_handler

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background() // This is a default context

var RedisClient *redis.Client

func init() {
	var redis_addr string
	if os.Getenv("IN_DOCKER") == "TRUE" {
		redis_addr = "redis:6379"
	} else {
		redis_addr = "localhost:6379"
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr: redis_addr,
	})
}
