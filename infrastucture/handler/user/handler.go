package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ninosistemas10/ecommerce/domain/user"
	"github.com/ninosistemas10/ecommerce/infrastucture/handler/response"
	"github.com/ninosistemas10/ecommerce/model"
)

type handler struct {
	useCase   user.UseCase
	responser response.API
}

func newHandler(uc user.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) Create(c *fiber.Ctx) error {
	m := model.User{}
	if err := c.BodyParser(&m); err != nil {
		return h.responser.BindFailed(err)
	}
	if err := h.useCase.Create(&m); err != nil {
		return h.responser.Error(c, "useCase.Create()", errors.New("Couldn't parse the ID"))
	}

	status, response := h.responser.Created(m)
	return c.Status(status).JSON(response)
}

func (h handler) MySelf(c *fiber.Ctx) error {
	ID, ok := c.Locals("userID").(uuid.UUID)
	if !ok {
		return h.responser.Error(c, "", errors.New("Couldn't parese the ID"))
	}

	u, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.responser.Error(c, "useCase.GetByID", err)
	}
	status, response := h.responser.OK(u)
	return c.Status(status).JSON(response)
}

func (h handler) GetAll(c *fiber.Ctx) error {
	users, err := h.useCase.GetAll()
	if err != nil {
		return h.responser.Error(c, "useCase.GetAll()", err)
	}

	status, response := h.responser.OK(users)
	return c.Status(status).JSON(response)
}
