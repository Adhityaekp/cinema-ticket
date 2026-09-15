package model

import (
	"time"

	"github.com/google/uuid"
)

type Refund struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	BookingID  uuid.UUID  `db:"booking_id" json:"booking_id"`
	RefundID   *string    `db:"refund_id" json:"refund_id,omitempty"`
	Amount     float64    `db:"amount" json:"amount"`
	Reason     string     `db:"reason" json:"reason"`
	Status     string     `db:"status" json:"status"`
	RefundedAt *time.Time `db:"refunded_at" json:"refunded_at,omitempty"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updated_at"`
}