package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kwanok/coupon-issue-project/coupon"
	"github.com/kwanok/coupon-issue-project/rdb"
)

var serverPort = flag.String("serverPort", ":3000", "api server serverPort")
var redisAddr = flag.String("redisAddr", "0.0.0.0:6379", "redis address")

func main() {
	flag.Parse()

	StartServer()
}

func StartServer() {
	redisClient := rdb.NewClient(*redisAddr)
	app := fiber.New()

	app.Use(logger.New(logger.ConfigDefault))
	coupon.Init(app, redisClient)

	err := app.Listen(*serverPort)
	if err != nil {
		panic(err)
	}
}
