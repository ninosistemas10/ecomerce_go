package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ninosistemas10/ecommerce/domain/user"
	"github.com/ninosistemas10/ecommerce/infrastructure/handler/middle"
	storageUser "github.com/ninosistemas10/ecommerce/infrastructure/postgres/user"
)

func NewRouter(app *fiber.App, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)
	authMiddleware := middle.New()

	// Convert the authMiddleware functions to the appropriate fiber.Handler type
	adminRoutes(app, h, authMiddleware.IsValid, authMiddleware.IsAdmin)
	publicRoutes(app, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	storage := storageUser.New(dbPool)
	useCase := user.New(storage)

	return newHandler(useCase)
}

func adminRoutes(app *fiber.App, h handler, middlewares ...fiber.Handler) {
	g := app.Group("/api/v1/admin/users", middlewares...)

	g.Get("", h.GetAll)
}

func publicRoutes(app *fiber.App, h handler) {
	g := app.Group("/api/v1/public/users")

	g.Post("", h.Create)
}
