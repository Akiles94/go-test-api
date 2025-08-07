package shared_ports

import (
	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
)

type JWTServicePort interface {
	Sign(claims *value_objects.UserClaims) (string, error)
	Parse(token string) (*value_objects.UserClaims, error)
	IsValid(token string) bool
}
