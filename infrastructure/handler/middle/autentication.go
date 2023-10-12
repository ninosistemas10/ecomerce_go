package middle

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gofiber/fiber/v2"

	"github.com/ninosistemas10/ecommerce/infrastructure/handler/response"
	"github.com/ninosistemas10/ecommerce/model"
)

type AuthMiddleware struct {
	responser response.API
}

// New crea una nueva instancia de AuthMiddleware
func New() AuthMiddleware {
	return AuthMiddleware{}
}

// IsValid es un middleware para validar el token de autenticación
func (am AuthMiddleware) IsValid(ctx *fiber.Ctx) error {
	// Obtener el token del encabezado de la solicitud
	token, err := getTokenFromRequest(ctx)
	if err != nil {
		return am.responser.BindFailed(err)
	}

	// Validar el token
	isValid, claims := am.validate(token)
	if !isValid {
		err := errors.New("el token no es válido")
		return am.responser.BindFailed(err)
	}

	// Establecer datos del usuario en el contexto de Fiber
	ctx.Locals("userID", claims.UserID)
	ctx.Locals("email", claims.Email)
	ctx.Locals("isAdmin", claims.IsAdmin)

	return ctx.Next()
}

// IsAdmin es un middleware para verificar si el usuario es un administrador
func (am AuthMiddleware) IsAdmin(ctx *fiber.Ctx) error {
	// Verificar si el usuario es un administrador
	isAdmin, ok := ctx.Locals("isAdmin").(bool)
	if !isAdmin || !ok {
		err := errors.New("no eres administrador")
		return am.responser.BindFailed(err)
	}

	return ctx.Next()
}

// validate valida un token JWT
func (am AuthMiddleware) validate(token string) (bool, model.JWTCustomClaims) {
	claims := &model.JWTCustomClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		log.Println(token)
		log.Println(os.Getenv("JWT_SECRET_KEY"))
		log.Println(err)
		return false, model.JWTCustomClaims{}
	}

	if !parsedToken.Valid {
		log.Println("Token no válido")
		return false, model.JWTCustomClaims{}
	}

	return true, *claims
}

// getTokenFromRequest obtiene el token del encabezado de la solicitud
func getTokenFromRequest(ctx *fiber.Ctx) (string, error) {
	data := ctx.Get("Authorization")
	if data == "" {
		return "", errors.New("el encabezado de autorización está vacío")
	}

	if strings.HasPrefix(data, "Bearer") {
		return strings.TrimSpace(data[7:]), nil
	}

	return data, nil
}
