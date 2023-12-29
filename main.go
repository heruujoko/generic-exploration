package main

import (
	"github.com/gofiber/fiber/v2"
	"gogenerics/common"
	"gogenerics/transport/http"
	"log"
	"time"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/ok", func(c *fiber.Ctx) error {
		return c.JSON(http.ResponseWithHTTP[common.HealthCheck](common.HealthCheck{
			Version:    "0.0.1",
			ServerTime: time.Now(),
		}, nil))
	})

	app.Get("/slice", func(c *fiber.Ctx) error {
		logs := []common.DummyLog{
			common.DummyLog{
				Action:     "SETUP",
				ServerTime: time.Now().Add(-5 * time.Second),
			},
			common.DummyLog{
				Action:     "RUN",
				ServerTime: time.Now(),
			},
		}

		return c.JSON(http.ResponseWithHTTP[[]common.DummyLog](logs, nil))
	})

	app.Get("/err", func(c *fiber.Ctx) error {
		errCode := "RATE_EXCEEDED"
		return c.JSON(http.ResponseWithHTTP[*common.DummyLog](nil, &errCode))
	})

	log.Fatal(app.Listen(":3000"))
}
