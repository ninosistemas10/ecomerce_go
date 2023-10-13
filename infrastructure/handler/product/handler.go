package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ninosistemas10/ecommerce/domain/product"
	"github.com/ninosistemas10/ecommerce/infrastructure/handler/response"
	"github.com/ninosistemas10/ecommerce/model"
)



type handler struct {
	useCase product.UseCase
	response response.API
}

func newhandler(useCase product.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) Create(ctx *fiber.Ctx) error {
	m := model.Product{}
	if err := ctx.BodyParser(&m); err != nil { return h.response.BindFailed(err) }

	if err := h.useCase.Create(&m); err != nil { return h.response.Error(ctx, "useCase.Create()", err) }

	status, resp := h.response.Created(m)
	return ctx.Status(status).JSON(resp)

}

func (h handler) Update(ctx *fiber.Ctx) error {
	m := model.Product{}

	if err := ctx.BodyParser(&m); err != nil { return h.response.BindFailed(err) }

	ID, err := uuid.Parse(ctx.Params("id"))
	if err != nil { return h.response.BindFailed(err) }
	m.ID = ID

	if err := h.useCase.Update(&m); err != nil { return h.response.Error(ctx, "h.useCase.Update()", err) }

	status, resp := h.response.Update(m)
	return ctx.Status(status).JSON(resp)
}

func (h handler) Delete(ctx *fiber.Ctx) error {
	ID, err := uuid.Parse(ctx.Params("id"))
	if err != nil { h.response.BindFailed(err) }

	err = h.useCase.Delete(ID)
	if err != nil { return h.response.Error(ctx, "useCase.Delete()", err) }

	status, resp := h.response.Delete(nil)
	return ctx.Status(status).JSON(resp)
}

func (h handler) GetByID(ctx *fiber.Ctx) error {
	ID, err := uuid.Parse(ctx.Params("id"))
	if err != nil { return h.response.Error(ctx, "uuid.Parse()", err) }

	productData, err := h.useCase.GetByID(ID)
	if err != nil { return h.response.Error(ctx, "useCase.GeByID", err) }

	status, resp := h.response.OK(productData)
	return ctx.Status(status).JSON(resp)
}

func (h handler) GetAll(ctx *fiber.Ctx) error {
	products, err := h.useCase.GetAll()
	if err != nil { return h.response.Error(ctx, "useCase.GetAll()", err) }
	status, resp := h.response.OK(products)
	return ctx.Status(status).JSON(resp)
}


// func (h handler) GetPaginatedData(ctx *fiber.Ctx) error {
// 	// Obtén los parámetros de la consulta (query params)
// 	limit := 10 // Valor predeterminado
// 	if limitStr := ctx.Query("limit"); limitStr != "" {
// 		limit, err := strconv.Atoi(limitStr)
// 		if err != nil || limit <= 0 {
// 		  // Manejo de errores o establecer un valor predeterminado
// 			limit = 10
// 		}
// 	}

// 	page := 1 // Valor predeterminado
// 	if pageStr := ctx.Query("page"); pageStr != "" {
// 		page, err := strconv.Atoi(pageStr)
// 		if err != nil || page <= 0 {
// 		  // Manejo de errores o establecer un valor predeterminado
// 			page = 1
// 		}
// 	}

// 	// Calcula el índice inicial para la paginación
// 	startIndex := (page - 1) * limit

// 	// Realiza la consulta con los parámetros de paginación
// 	products, err := h.useCase.GetPaginatedData(startIndex, limit)
// 	if err != nil {
// 		return h.response.Error(ctx, "useCase.GetPaginatedData()", err)
// 	}

// 	// Puedes agregar lógica adicional aquí, como contar el número total de elementos, etc.

// 	// Preparar la respuesta JSON
// 	status, resp := h.response.OK(struct {
// 		Products []model.Product `json:"products"`
// 	}{
// 		Products: products,
// 	})
// 	return ctx.Status(status).JSON(resp)
// }
