package rdb

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var Client *redis.Client
var Ctx = context.Background()

func NewClient(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    "", // no password set
		DB:          0,  // use default DB
		ReadTimeout: 1 * time.Minute,
	})
}
