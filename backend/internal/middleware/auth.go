package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/breodoyo/niavo/backend/internal/auth"
	"github.com/breodoyo/niavo/backend/internal/common"
)

type contextKey string

const UserIDKey contextKey = "userID"

func RequireAuth(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				common.WriteError(w, http.StatusUnauthorized, "unauthorized", "missing authorization header")
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				common.WriteError(w, http.StatusUnauthorized, "unauthorized", "invalid authorization header format")
				return
			}

			tokenString := parts[1]

			claims, err := auth.ValidateToken(tokenString, jwtSecret)
			if err != nil {
				common.WriteError(w, http.StatusUnauthorized, "unauthorized", "invalid or expired token")
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}