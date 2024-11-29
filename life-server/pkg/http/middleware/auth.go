package middleware

import (
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
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

		next.ServeHTTP(w, r)
	})
}

func parseErr(w http.ResponseWriter) {
	http.Error(w, "Error parsing authentication token", http.StatusUnauthorized)
}
