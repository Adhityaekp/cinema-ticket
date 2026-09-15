package middleware

import (
	"context"
	"net/http"

	"github.com/Adhityaekp/cinema-ticket/pkg/jwt"
	"github.com/Adhityaekp/cinema-ticket/pkg/response"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
	RoleKey   contextKey = "role"
)

func AuthMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Ambil access token dari cookie
			cookie, err := r.Cookie("access_token")
			if err != nil {
				response.Error(
					w,
					http.StatusUnauthorized,
					"unauthorized",
					"access token tidak ditemukan",
				)
				return
			}

			// Validasi JWT
			claims, err := jwt.ValidateToken(
				cookie.Value,
				jwtSecret,
			)
			if err != nil {
				response.Error(
					w,
					http.StatusUnauthorized,
					"unauthorized",
					"access token tidak valid atau sudah expired",
				)
				return
			}

			// Pastikan token adalah access token
			if claims.Type != "access" {
				response.Error(
					w,
					http.StatusUnauthorized,
					"unauthorized",
					"token bukan access token",
				)
				return
			}

			// Simpan informasi user ke context
			ctx := context.WithValue(
				r.Context(),
				UserIDKey,
				claims.UserID,
			)

			ctx = context.WithValue(
				ctx,
				RoleKey,
				claims.Role,
			)

			// Lanjut ke handler berikutnya
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
