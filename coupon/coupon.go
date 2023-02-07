package coupon

import (
	"github.com/kwanok/coupon-issue-project/rdb"
	"github.com/redis/go-redis/v9"
	"log"
)

var queue RedisClient

type Coupon struct {
	Code string
}

type RedisClient struct {
	*redis.Client
}

func (c *RedisClient) push(coupon *Coupon) error {
	_, err := c.TxPipelined(rdb.Ctx, func(pipe redis.Pipeliner) error {
		c.RPush(rdb.Ctx, "coupon_codes", coupon.Code)
		return nil
	})

	if err != nil {
		log.Printf("Error pushing coupon to Redis: %v", err)
		return err
	}

	return nil
}

func (c *RedisClient) pop() (*Coupon, error) {
	code, err := c.LPop(rdb.Ctx, "coupon_codes").Result()

	if err != nil {
		log.Printf("Error pop coupon from Redis: %v", err)
		return nil, err
	}

	return &Coupon{
		Code: code,
	}, nil
}
