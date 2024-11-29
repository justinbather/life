package middleware

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			fmt.Println("Unable to parse auth token")
			http.Error(w, "Unable to parse token", http.StatusForbidden)
			return
		}

		if token != "Bearer some-token" {
			fmt.Println("Unknown auth token")
			http.Error(w, "Forbidden", http.StatusForbidden)
		}

		next.ServeHTTP(w, r)
	})
}
