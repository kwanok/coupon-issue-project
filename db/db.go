package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var Ctx = context.Background()
var Client = NewRedisClient()

func NewRedisClient() *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		ReadTimeout: time.Duration(60) * time.Second,
		PoolSize:    1000,
	})

	err := r.Ping(Ctx).Err()
	if err != nil {
		panic(err)
	}

	return r
}
