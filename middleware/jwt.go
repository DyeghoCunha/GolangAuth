package middleware

import (
	"context"
	"github.com/dyeghocunha/golang-auth/repository"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func JWTAuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token ausente", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["email"] == nil {
			http.Error(w, "Claims inválidos", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_email", claims["email"].(string))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func JWTAuthMiddlewareWith2FA(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token ausente", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["email"] == nil {
			http.Error(w, "Claims inválidos", http.StatusUnauthorized)
			return
		}
		email := claims["email"].(string)

		user, err := repository.GetUserByEmail(email)
		if err != nil || !user.IsTwoFAEnabled {
			http.Error(w, "2FA não está habilitado", http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), "user_email", email)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
