package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/justinbather/life/life-server/pkg/service"
	"github.com/justinbather/prettylog"
)

type ctxKey string

const UserCtxKey ctxKey = "user"

type Middleware interface {
	Protect(next http.Handler) http.Handler
	Recoverer(next http.Handler) http.Handler
	Tracer(next http.Handler) http.Handler
}

func NewMiddleware(authService service.AuthService, userService service.UserService, logger *prettylog.Logger) Middleware {
	return &middleware{authService: authService, userService: userService, logger: logger}
}

type middleware struct {
	authService service.AuthService
	userService service.UserService
	logger      *prettylog.Logger
}

// Authentication Middleware
// Pulls and verifys the Bearer token from request header
func (m *middleware) Protect(next http.Handler) http.Handler {
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

		// TODO: optimizable
		_, err = m.userService.GetUserById(r.Context(), userId)
		if err != nil {
			authErr(w)
			return
		}

		ctx := context.WithValue(r.Context(), UserCtxKey, userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Recovers any panics and returns a 500 error code to the user
func (m *middleware) Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				m.logger.Errorf("Recovery middleware caught a panic. %v", rec)
				http.Error(w, "Internal Server Error.", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (m *middleware) Tracer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.logger.Infof("Starting Request")
		start := time.Now()
		next.ServeHTTP(w, r)
		m.logger.Infof("Completed request in %f seconds", time.Since(start).Seconds())

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
	http.Error(w, "Error parsing authentication token", http.StatusBadRequest)
}

func authErr(w http.ResponseWriter) {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
