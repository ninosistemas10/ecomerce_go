package main

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func newHTTP(errorHandler func(*fiber.Ctx, error)) *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	// Obtén los valores de entorno y conviértelos a slices de strings
	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	allowedOriginsString := strings.Join(allowedOrigins, ",")

	corsConfig := cors.New(cors.Config{
      	AllowOrigins: allowedOriginsString,  // Convierte el slice a una cadena
      	AllowMethods: strings.Join(strings.Split(os.Getenv("ALLOWED_METHODS"), ","), ","),
    	})

	app.Use(corsConfig)

	app.Use(errorHandler)

	return app
}
