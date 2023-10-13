package purchaseorder

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	purchaseorder "github.com/ninosistemas10/ecommerce/domain/purchaseOrder"
	"github.com/ninosistemas10/ecommerce/infrastructure/handler/middle"
	purchaseorderStorage "github.com/ninosistemas10/ecommerce/infrastructure/postgres/purchaseorder"
)

func NewRouter(app *fiber.App, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middle.New()
	privateRoutes(app, h ,authMiddleware.IsValid)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCase := purchaseorder.New(purchaseorderStorage.New(dbPool))
	return NewHandler(useCase)

}

// privateRoutes handle the routes that require a token
func privateRoutes(app *fiber.App, h handler, middlewares ...func(*fiber.Ctx) error) {
	route := app.Group("/api/v1/private/purchase-orders", middlewares...)

	route.Post("", h.Create)
}






