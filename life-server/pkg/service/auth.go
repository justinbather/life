package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret_key = []byte("secret-key")

type AuthService interface {
	Authenticate(token string) (string, error)
	CreateToken(id string) (string, time.Time, error)
}

func NewAuthService() AuthService {
	return &authService{}
}

type authService struct{}

func (s *authService) Authenticate(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret_key, nil
	})
	if err != nil {
		return "", nil
	}

	if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
		return "", fmt.Errorf("Invalid Algorithm")
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		if claims.Expires.Before(time.Now()) {
			return "", fmt.Errorf("Expired token. Reauthenticate")
		}

		return claims.UserId, nil
	} else {
		return "", fmt.Errorf("Invalid token")
	}
}

type CustomClaims struct {
	UserId  string    `json:"userId"`
	Expires time.Time `json:"expires"`
	jwt.RegisteredClaims
}

func (s *authService) CreateToken(id string) (string, time.Time, error) {
	expires := time.Now().Add(24 * time.Hour)

	claims := CustomClaims{
		UserId:  id,
		Expires: expires,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(secret_key)
	if err != nil {
		return "", time.Time{}, err
	}
	return signed, expires, nil
}
