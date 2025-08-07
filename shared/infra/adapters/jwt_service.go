package shared_adapters

import (
	"errors"

	"github.com/Akiles94/go-test-api/config"
	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey []byte
}

func NewJWTService() shared_ports.JWTServicePort {
	return &JWTService{
		secretKey: []byte(config.Env.JWTSecret),
	}
}

func (j *JWTService) Sign(claims *value_objects.UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

func (j *JWTService) Parse(tokenString string) (*value_objects.UserClaims, error) {
	claims := &value_objects.UserClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (j *JWTService) IsValid(tokenString string) bool {
	_, err := j.Parse(tokenString)
	return err == nil
}
