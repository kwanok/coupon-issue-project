package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/kwanok/coupon-generate-study/coupon"
	"github.com/kwanok/coupon-generate-study/db"
	"log"
)

func main() {
	db.NewRedisClient()

	app := fiber.New(fiber.Config{
		IdleTimeout: 10 * 1000,
	})

	app.Get("/getCoupon", func(c *fiber.Ctx) error {
		popCoupon := coupon.PopCoupon(context.Background())

		return c.SendString(popCoupon)
	})

	log.Fatal(app.Listen(":3000"))
}
