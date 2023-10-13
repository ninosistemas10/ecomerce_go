package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/ninosistemas10/ecommerce/model"
)

const (
	BindFailed      = "bind_failed"
	Ok              = "ok"
	RecordCreated   = "record_created"
	RecordUpdated   = "record_updated"
	RecordDeleted   = "record_deleted"
	UnExpectederror = "unexpected_error"
	AuthError       = "authorization_error"
)

type API struct{}

func New() API {
	return API{}
}

func (a API) OK(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:    data,
		Messages: model.Responses{{Code: Ok, Message: "!listo¡"}},
	}
}

func (a API) Created(data interface{}) (int, model.MessageResponse) {
	return http.StatusCreated, model.MessageResponse{
		Data:    data,
		Messages: model.Responses{{Code: RecordCreated, Message: "!listo¡"}},
	}
}

func (a API) Update(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:    data,
		Messages: model.Responses{{Code: RecordUpdated, Message: "!listo¡"}},
	}
}

func (a API) Delete(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:    data,
		Messages: model.Responses{{Code: RecordDeleted, Message: "!listo¡"}},
	}
}

func (a API) BindFailed(err error) error {
	e := model.NewError()
	e.Err = err
	e.Code = BindFailed
	e.StatusHTTP = http.StatusBadRequest
	e.Who = "c.Bind()"

	log.Warnf("%s", e.Error())
	return &e
}

func (a API) Error(c *fiber.Ctx, who string, err error) *model.ErrorE {
	e := model.NewError()
	e.Err = err
	e.APIMessage = "¡Uy! metimos la pata, disculpanos lo solucionaremos"
	e.Code = UnExpectederror
	e.StatusHTTP = http.StatusInternalServerError
	e.Who = who

	userID, ok := c.Locals("userID").(uuid.UUID)
	// Solo para evitar el error de pánico
	if !ok {
		log.Errorf("no se puede obtener/analizar UUID desde userID")
	}
	e.UserID = userID.String()

	log.Errorf("%s", e.Error())
	return &e
}
