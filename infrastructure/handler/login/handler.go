package login

import (
	"database/sql"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/ninosistemas10/ecommerce/domain/login"
	"github.com/ninosistemas10/ecommerce/infrastructure/handler/response"
	"github.com/ninosistemas10/ecommerce/model"
)

//hola como estan
type handler struct {
	useCase   login.UseCase
	responser response.API
}

func newHandler(useCase login.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) Login(c *fiber.Ctx) error {
	m := new(model.Login)
	if err := c.BodyParser(m); err != nil {
		return h.responser.BindFailed(err)
	}

	u, t, err := h.useCase.Login(m.Email, m.Password, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		if strings.Contains(err.Error(), "bcrypt.CompareHashAndPassword()") || errors.Is(err, sql.ErrNoRows) {
			resp := model.MessageResponse{
				Data:     "wrong user or password",
				Messages: model.Responses{{Code: response.AuthError, Message: "wrong user or password"}},
			}
			return c.Status(http.StatusBadRequest).JSON(resp)
		}
		return h.responser.Error(c, "useCase.Login()", err)
	}

	httpStatus, resp := h.responser.OK(map[string]interface{}{"user": u, "token": t})
	return c.Status(httpStatus).JSON(resp)

}
