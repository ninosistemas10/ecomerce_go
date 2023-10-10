package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2/log"
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
		Message: model.Responses{{Code: Ok, Message: "¡listo"}},
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
