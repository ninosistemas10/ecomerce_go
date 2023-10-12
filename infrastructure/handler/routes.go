package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ninosistemas10/ecommerce/infrastructure/handler/login"
	"github.com/ninosistemas10/ecommerce/infrastructure/handler/user"
)



func InitRoutes(app *fiber.App, dbPool *pgxpool.Pool) {
	health(app)

	// A
	// B
	// C

	// I


	// L
	login.NewRouter(app, dbPool)

	// P


	// U
	user.NewRouter(app, dbPool)
}

func health(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"time":         time.Now().String(),
			"message":      "Hello World!",
			"service_name": "",
		})
	})
}
