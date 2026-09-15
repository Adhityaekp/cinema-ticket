package dto

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email,max=150"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type VerifyEmailRequest struct {
	Token string `json:"token" validate:"required"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}
