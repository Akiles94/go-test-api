package value_objects

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

type UserClaims struct {
	UserID    uuid.UUID `json:"user_id"`
	Email     *string   `json:"email"`
	Name      *string   `json:"name"`
	Role      string    `json:"role"`
	TokenType TokenType `json:"token_type"`
	jwt.RegisteredClaims
}
