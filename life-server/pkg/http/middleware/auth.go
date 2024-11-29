package middleware

import (
	"context"
	"fmt"
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
			return
		}

		if tokenArr[0] != "Bearer" {
			parseErr(w)
			return
		}

		userId, err := m.authService.Authenticate(tokenArr[1])
		if err != nil {
			authErr(w)
			return
		}

		fmt.Printf("Authenticated user: %s\n", userId)

		ctx := context.WithValue(r.Context(), UserCtxKey, userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserFromCtx(ctx context.Context) (string, error) {
	user := ctx.Value(UserCtxKey)
	if user == nil {
		return "", fmt.Errorf("Error getting user from context")
	}

	fmt.Printf("User from context: %v\n", user)

	return user.(string), nil
}

func parseErr(w http.ResponseWriter) {
	http.Error(w, "Error parsing authentication token", http.StatusUnauthorized)
}

func authErr(w http.ResponseWriter) {
	http.Error(w, "Error authenticating", http.StatusUnauthorized)
}
