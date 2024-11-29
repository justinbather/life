package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret_key = []byte("secret-key")

type AuthService interface {
	Authenticate(token string) (string, error)
	CreateToken(id string) (string, int64, error)
}

func NewAuthService() AuthService {
	return &authService{}
}

type authService struct{}

func (s *authService) Authenticate(token string) (string, error) {
	return "", nil
}

func (s *authService) CreateToken(id string) (string, int64, error) {
	expires := time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":      id,
			"expires": expires,
		})

	signed, err := token.SignedString(secret_key)
	if err != nil {
		return "", 0, err
	}
	return signed, expires, nil
}
