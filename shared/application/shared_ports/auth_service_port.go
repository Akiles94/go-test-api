package shared_ports

import (
	"context"

	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
	"github.com/google/uuid"
)

type AuthServicePort interface {
	GenerateToken(ctx context.Context, userID uuid.UUID, email, name *string) (*value_objects.JWTToken, error)
	ValidateToken(ctx context.Context, token string) (*value_objects.UserClaims, error)
	RefreshToken(ctx context.Context, refreshToken string) (*value_objects.JWTToken, error)
	GenerateRefreshToken(ctx context.Context, userID uuid.UUID, email string) (*value_objects.JWTToken, error)
}
