package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/justinbather/life/life-server/pkg/service"
)

type ctxKey string

const UserCtxKey ctxKey = "user"

type AuthMiddleware interface {
	Protect(next http.Handler) http.Handler
}

func NewAuthMiddleware(authService service.AuthService) AuthMiddleware {
	return &authMiddleware{authService: authService}
}

type authMiddleware struct {
	authService service.AuthService
}

// Authentication Middleware
// Pulls and verifys the Bearer token from request header
func (m *authMiddleware) Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt, ok := parseHeader(r.Header)
		if !ok {
			parseErr(w)
			return
		}

		userId, err := m.authService.Authenticate(jwt)
		if err != nil {
			authErr(w)
			return
		}

		ctx := context.WithValue(r.Context(), UserCtxKey, userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func parseHeader(header http.Header) (string, bool) {
	token := header.Get("Authorization")
	if token == "" {
		return "", false
	}

	tokenArr := strings.Split(token, " ")

	if len(tokenArr) != 2 {
		return "", false
	}

	if tokenArr[0] != "Bearer" {
		return "", false
	}

	return tokenArr[1], true
}

func parseErr(w http.ResponseWriter) {
	http.Error(w, "Error parsing authentication token", http.StatusUnauthorized)
}

func authErr(w http.ResponseWriter) {
	http.Error(w, "Error authenticating", http.StatusUnauthorized)
}
