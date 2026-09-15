package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Adhityaekp/cinema-ticket/config"
	"github.com/Adhityaekp/cinema-ticket/internal/dto"
	"github.com/Adhityaekp/cinema-ticket/internal/service"
	jwtpkg "github.com/Adhityaekp/cinema-ticket/pkg/jwt"
	"github.com/Adhityaekp/cinema-ticket/pkg/response"
)

type AuthHandler struct {
	authService *service.AuthService
	cfg         *config.Config
}

func NewAuthHandler(
	authService *service.AuthService,
	cfg *config.Config,
) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		cfg:         cfg,
	}
}

// Register godoc
// @Summary Register customer
// @Description Membuat akun customer baru dan mengirim email verifikasi.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /auth/register [post]
func (h *AuthHandler) Register(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"Request tidak valid",
			nil,
		)
		return
	}

	user, err := h.authService.Register(
		r.Context(),
		req,
	)

	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
			nil,
		)
		return
	}

	response.Created(
		w,
		"Registrasi berhasil. Silakan cek email untuk verifikasi.",
		map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	)
}

// VerifyEmail godoc
// @Summary Verify email
// @Description Memverifikasi email customer menggunakan verification token.
// @Tags Auth
// @Produce json
// @Param token query string true "Verification token"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /auth/verify-email [get]
func (h *AuthHandler) VerifyEmail(
	w http.ResponseWriter,
	r *http.Request,
) {
	token := r.URL.Query().Get("token")

	if token == "" {
		response.Error(
			w,
			http.StatusBadRequest,
			"Token wajib diisi",
			nil,
		)
		return
	}

	err := h.authService.VerifyEmail(
		r.Context(),
		token,
	)

	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
			nil,
		)
		return
	}

	response.Success(
		w,
		"Email berhasil diverifikasi. Silakan login.",
		nil,
	)
}

// Login godoc
// @Summary Login
// @Description Login menggunakan email dan password.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login request"
// @Success 200 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /auth/login [post]
func (h *AuthHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"Request tidak valid",
			nil,
		)
		return
	}

	user, err := h.authService.Login(
		r.Context(),
		req,
	)

	if err != nil {
		response.Error(
			w,
			http.StatusUnauthorized,
			err.Error(),
			nil,
		)
		return
	}

	accessToken, err := jwtpkg.GenerateAccessToken(
		user.ID,
		user.Role,
		h.cfg.JWTSecret,
		h.cfg.JWTAccessExpireMinutes,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"Gagal membuat access token",
			nil,
		)
		return
	}

	refreshToken, err := jwtpkg.GenerateRefreshToken(
		user.ID,
		user.Role,
		h.cfg.JWTSecret,
		h.cfg.JWTRefreshExpireDays,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"Gagal membuat refresh token",
			nil,
		)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   h.cfg.CookieSecure,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.JWTAccessExpireMinutes * 60,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/api/auth",
		HttpOnly: true,
		Secure:   h.cfg.CookieSecure,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.JWTRefreshExpireDays * 24 * 60 * 60,
	})

	response.Success(
		w,
		"Login berhasil",
		map[string]interface{}{
			"user": map[string]interface{}{
				"id":        user.ID,
				"name":      user.Name,
				"email":     user.Email,
				"role":      user.Role,
				"is_active": user.IsActive,
			},
		},
	)
}
