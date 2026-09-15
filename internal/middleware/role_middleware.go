package middleware

import (
	"net/http"

	"github.com/Adhityaekp/cinema-ticket/pkg/response"
)

func RoleMiddleware(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			role := r.Context().Value(RoleKey)

			if role == nil {
				response.Error(
					w,
					http.StatusUnauthorized,
					"unauthorized",
					"role user tidak ditemukan",
				)
				return
			}

			userRole, ok := role.(string)
			if !ok {
				response.Error(
					w,
					http.StatusUnauthorized,
					"unauthorized",
					"role user tidak valid",
				)
				return
			}

			if userRole != requiredRole {
				response.Error(
					w,
					http.StatusForbidden,
					"forbidden",
					"akses hanya untuk "+requiredRole,
				)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
