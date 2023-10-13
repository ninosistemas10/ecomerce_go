package login

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ninosistemas10/ecommerce/domain/login"
	"github.com/ninosistemas10/ecommerce/domain/user"

	userStorage "github.com/ninosistemas10/ecommerce/infrastructure/postgres/user"
)

// NewRouter returns a router to handle model.Login requests
func NewRouter(app *fiber.App, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	// build middlewares to validate permissions on the routes

	publicRoutes(app, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCaseUser := user.New(userStorage.New(dbPool))
	useCase := login.New(useCaseUser)
	return newHandler(useCase)
}


// publicRoutes handle the routes that not requires a validation of any kind to be use
func publicRoutes(app *fiber.App, h handler) {
	route := app.Group("/api/v1/public/login")

	route.Post("", h.Login)
}
