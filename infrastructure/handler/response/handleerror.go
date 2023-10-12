package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ninosistemas10/ecommerce/model"
)

// HTTPErrorHandler es un manejador de errores HTTP
func HTTPErrorHandler(ctx *fiber.Ctx, err error){
	// error personalizado
	e, ok := err.(*model.ErrorE)
	if ok {
		_ = ctx.JSON(getResponseError(e))
		return
	}

	// verificar error de echo
	if fiberErr, ok := err.(*fiber.Error); ok {
		msg := fiberErr.Message
		if msg == "" {
			msg = "¡Upps! algo inesperado ocurrió"
		}

		_ = ctx.Status(fiberErr.Code).JSON( model.MessageResponse {
			Errors: model.Responses{
				{Code: UnExpectederror , Message: msg},
			},
		})
	return
	}
	_ = ctx.Status(http.StatusInternalServerError).JSON(model.MessageResponse{
		Errors: model.Responses{
			{Code: UnExpectederror, Message: ""},
		},
	})
}

	// si el controlador no devuelve un "model.Error", entonces devuelve una respuesta JSON de error genérica



func getResponseError(err *model.ErrorE) fiber.Map {
	output := fiber.Map{}
	if !err.HasCode() {
		err.Code = UnExpectederror
	}

	if err.HasData() {
		output["data"] = err.Data
	}

	if !err.HasStatusHttp() {
		err.StatusHTTP = http.StatusInternalServerError
	}

	output["status"] = err.StatusHTTP
	output["errors"] = []model.Response{
		{Code: err.Code, Message: err.APIMessage},
	}

	return output
}
