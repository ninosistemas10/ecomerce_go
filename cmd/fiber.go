package main

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func newHTTP() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// tu manejador de errores aquí
			return ctx.Status(500).SendString(err.Error())
		},
	})

	app.Use(logger.New())
	app.Use(recover.New())

	corsConfig := cors.Config{
		AllowOrigins: strings.Join(strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","), ","),
		AllowMethods: strings.Join(strings.Split(os.Getenv("ALLOWED_METHODS"), ","), ","),
	}

	app.Use(cors.New(corsConfig))

	return app
}
