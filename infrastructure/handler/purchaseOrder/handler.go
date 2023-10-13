package purchaseorder

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	purchaseorder "github.com/ninosistemas10/ecommerce/domain/purchaseOrder"
	"github.com/ninosistemas10/ecommerce/infrastructure/handler/response"
	"github.com/ninosistemas10/ecommerce/model"
)

type handler struct {
	useCase purchaseorder.UseCase
	response response.API
}

func NewHandler(useCase purchaseorder.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) Create (c *fiber.Ctx) error {
	m := model.PurchaseOrder{}
	if err := c.BodyParser(&m); err != nil { return h.response.BindFailed(err) }

	UserID, ok := c.Locals("userID").(uuid.UUID)
	if !ok { return h.response.Error(c, "c.Local.(uuid.UUID)", errors.New("Can't parse uuid")) }
	m.UserID = UserID

	if err := h.useCase.Create(&m); err != nil { return h.response.Error(c, "useCase.Create()", err) }

	status, resp := h.response.Created(m)
	return c.Status(status).JSON(resp)
}

