package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTCustomClaims struct {
	UserID  uuid.UUID `json:"user_id"`
	Email   string    `json:"email"`
	IsAdmin bool      `json:"is_admin"`
	jwt.StandardClaims
}
