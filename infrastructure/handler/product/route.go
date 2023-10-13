package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ninosistemas10/ecommerce/domain/product"
	"github.com/ninosistemas10/ecommerce/infrastructure/handler/middle"
	productStorage "github.com/ninosistemas10/ecommerce/infrastructure/postgres/product"
)


func NewRouter(app *fiber.App, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middle.New()

	adminRoutes(app, h, authMiddleware.IsValid, authMiddleware.IsAdmin)
	publicRoutes(app, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCase := product.New(productStorage.New(dbPool))
	return newhandler(useCase)
}

func adminRoutes(app *fiber.App, h handler, middlewares ...func(*fiber.Ctx) error) {
	route := app.Group("/api/v1/admin/products", middlewares...)

	route.Post("", h.Create)
	route.Put("/:id", h.Update)
	route.Delete("/:id", h.Delete)
	route.Get("", h.GetAll)
	route.Get("/:id", h.GetByID)
}

func publicRoutes(app *fiber.App, h handler) {
	route := app.Group("/api/v1/public/products")

	route.Get("", h.GetAll)
	route.Get("/:id", h.GetByID)
}
