package model

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	BookingCode string     `db:"booking_code" json:"booking_code"`
	UserID      uuid.UUID  `db:"user_id" json:"user_id"`
	ShowtimeID  uuid.UUID  `db:"showtime_id" json:"showtime_id"`
	TotalAmount float64    `db:"total_amount" json:"total_amount"`
	Status      string     `db:"status" json:"status"`
	ExpiresAt   *time.Time `db:"expires_at" json:"expires_at,omitempty"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
}