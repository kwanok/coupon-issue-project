package coupon

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func Init(app *fiber.App, redis *redis.Client) {
	queue = RedisClient{redis}

	coupons := app.Group("/coupons")
	{
		coupons.Get("/", Send)
	}
}

func Send(c *fiber.Ctx) error {
	cp, err := queue.pop()

	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Server Down")
	}

	return c.JSON(cp)
}
