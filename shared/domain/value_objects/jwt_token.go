package value_objects

import (
	"time"
)

type JWTToken struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn   int64     `json:"expires_in"`
	ExpiresAt   time.Time `json:"expires_at"`
}

func NewJWTToken(accessToken string, expiresIn int64) *JWTToken {
	return &JWTToken{
		AccessToken: accessToken,
		ExpiresIn:   expiresIn,
		ExpiresAt:   time.Now().Add(time.Duration(expiresIn) * time.Second),
	}
}
