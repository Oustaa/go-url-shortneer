package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/oustaa/go-url-shortner/internal/utils"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		if bearerToken == "" {
			http.Error(w, "bearer token is not provided", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(bearerToken, "Bearer ")

		if len(parts) != 2 {
			http.Error(w, "invalid authorization format", http.StatusUnauthorized)
			return
		}

		token := parts[1]

		claimed, err := utils.ValidateToken(token)
		if err != nil {
			http.Error(w, fmt.Sprintf("bearer token is invalid, %s.", err.Error()), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), utils.UserIDKey, claimed.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
