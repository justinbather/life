package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/justinbather/prettylog"
)

var secret_key = []byte("secret-key")

type AuthService interface {
	Authenticate(ctx context.Context, token string) (string, error)
	CreateToken(id string) (string, time.Time, error)
}

func NewAuthService(db *pgx.Conn, userService UserService, cacheWindow int, logger *prettylog.Logger) AuthService {
	return &authService{db, userService, make(map[string]time.Time), &sync.Mutex{}, cacheWindow, logger}
}

// TODO: pull this into a caching service
type authService struct {
	db          *pgx.Conn
	userService UserService
	cache       map[string]time.Time
	m           *sync.Mutex
	cacheWindow int
	logger      *prettylog.Logger
}

func (s *authService) inCache(user string) (time.Time, bool) {
	s.m.Lock()
	val, ok := s.cache[user]
	s.m.Unlock()

	if !ok {
		s.logger.Infof("Cache miss")
		return time.Time{}, false
	}

	s.logger.Infof("Cache hit")
	return val, true
}

func (s *authService) revalidateCache(ctx context.Context, user string) error {
	_, err := s.userService.GetUserById(ctx, user)
	if err != nil {
		return err
	}

	s.updateCache(user)

	s.logger.Infof("Cache revalidated")
	return nil
}

func (s *authService) updateCache(user string) {
	s.m.Lock()
	s.cache[user] = time.Now()
	s.m.Unlock()

	s.logger.Infof("Cache updated")
}

func (s *authService) Authenticate(ctx context.Context, tokenString string) (string, error) {
	token, err := parseToken(tokenString)
	if err != nil {
		return "", fmt.Errorf("Error parsing authentication token")
	}

	ok := validateToken(token)
	if !ok {
		return "", fmt.Errorf("Invalid or malformed authentication token")
	}

	userId, err := extractUser(token)
	if err != nil {
		return "", err
	}

	timeAdded, exists := s.inCache(userId)
	if !exists {
		err := s.revalidateCache(ctx, userId)
		if err != nil {
			return "", err
		}

		return userId, nil
	}

	// invalid after 1 hour
	if timeAdded.Before(time.Now().Add(-time.Duration(s.cacheWindow) * time.Hour)) {
		err := s.revalidateCache(ctx, userId)
		if err != nil {
			return "", err
		}

		s.logger.Infof("Cache hit, but invalid")
		return userId, nil
	}

	return userId, nil
}

func (s *authService) CreateToken(id string) (string, time.Time, error) {
	expires := time.Now().Add(24 * time.Hour)

	claims := CustomClaims{
		UserId:  id,
		Expires: expires,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires),
			Issuer:    "life-server",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(secret_key)
	if err != nil {
		return "", time.Time{}, err
	}

	return signed, expires, nil
}

type CustomClaims struct {
	UserId  string    `json:"userId"`
	Expires time.Time `json:"expires"`
	jwt.RegisteredClaims
}

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret_key, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func validateToken(token *jwt.Token) bool {
	if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
		return false
	}

	return true
}

func extractUser(token *jwt.Token) (string, error) {
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		if claims.RegisteredClaims.ExpiresAt.Before(time.Now()) {
			return "", fmt.Errorf("Expired token. Reauthenticate")
		}

		return claims.UserId, nil
	} else {

		return "", fmt.Errorf("Invalid token")
	}
}
