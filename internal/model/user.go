package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                         uuid.UUID  `db:"id" json:"id"`
	Name                       string     `db:"name" json:"name"`
	Email                      string     `db:"email" json:"email"`
	Password                   string     `db:"password" json:"-"`
	Role                       string     `db:"role" json:"role"`
	IsActive                   bool       `db:"is_active" json:"is_active"`
	EmailVerifiedAt            *time.Time `db:"email_verified_at" json:"email_verified_at,omitempty"`
	EmailVerificationToken     *string    `db:"email_verification_token" json:"-"`
	EmailVerificationExpiresAt *time.Time `db:"email_verification_expires_at" json:"-"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
