package shared_adapters

import (
	"context"
	"errors"
	"time"

	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const oneDayInHours = 24

type AuthService struct {
	jwtService               shared_ports.JWTServicePort
	JWTRefreshExpirationDays int
}

func NewAuthService(jwtService shared_ports.JWTServicePort) shared_ports.AuthServicePort {
	return &AuthService{
		jwtService: jwtService,
	}
}

func (a *AuthService) GenerateToken(ctx context.Context, userID uuid.UUID, email, name *string) (*value_objects.JWTToken, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &value_objects.UserClaims{
		UserID:    userID,
		Email:     email,
		Name:      name,
		TokenType: value_objects.AccessToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-test-api",
		},
	}

	tokenString, err := a.jwtService.Sign(claims)
	if err != nil {
		return nil, err
	}

	return value_objects.NewJWTToken(tokenString, int64(24*60*60)), nil
}

func (a *AuthService) ValidateToken(ctx context.Context, token string) (*value_objects.UserClaims, error) {
	return a.jwtService.Parse(token)
}
func (a *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*value_objects.JWTToken, error) {
	claims, err := a.jwtService.Parse(refreshToken)
	if err != nil {
		return nil, err
	}
	if claims.TokenType != value_objects.RefreshToken {
		return nil, errors.New("invalid token type, expected refresh token")
	}
	if time.Now().After(claims.ExpiresAt.Time) {
		return nil, errors.New("refresh token expired")
	}
	return a.GenerateToken(ctx, claims.UserID, nil, nil)
}

func (a *AuthService) GenerateRefreshToken(ctx context.Context, userID uuid.UUID, email string) (*value_objects.JWTToken, error) {
	expirationTime := time.Now().Add(time.Duration(a.JWTRefreshExpirationDays) * oneDayInHours * time.Hour)

	claims := &value_objects.UserClaims{
		UserID:    userID,
		Email:     nil,
		Name:      nil,
		TokenType: value_objects.RefreshToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-test-api",
		},
	}

	tokenString, err := a.jwtService.Sign(claims)
	if err != nil {
		return nil, err
	}

	return value_objects.NewJWTToken(tokenString, int64(7*24*60*60)), nil
}
