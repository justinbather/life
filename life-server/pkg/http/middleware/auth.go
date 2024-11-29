package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/justinbather/life/life-server/pkg/service"
)

type AuthMiddleware interface {
	Protect(next http.Handler) http.Handler
}

func NewAuthMiddleware(authService service.AuthService) AuthMiddleware {
	return &authMiddleware{authService: authService}
}

type authMiddleware struct {
	authService service.AuthService
}

func (m *authMiddleware) Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			parseErr(w)
			return
		}

		tokenArr := strings.Split(token, " ")

		if len(tokenArr) != 2 {
			parseErr(w)
		}

		if tokenArr[0] != "Bearer" {
			parseErr(w)
		}

		userId, err := m.authService.Authenticate(tokenArr[1])
		if err != nil {
			authErr(w)
		}

		r.WithContext(context.WithValue(r.Context(), "userId", userId))

		next.ServeHTTP(w, r)
	})
}

func parseErr(w http.ResponseWriter) {
	http.Error(w, "Error parsing authentication token", http.StatusUnauthorized)
}

func authErr(w http.ResponseWriter) {
	http.Error(w, "Error authenticating", http.StatusUnauthorized)
}
