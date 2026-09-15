package model

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID            uuid.UUID  `db:"id" json:"id"`
	BookingID     uuid.UUID  `db:"booking_id" json:"booking_id"`
	PaymentMethod string     `db:"payment_method" json:"payment_method"`
	TransactionID *string    `db:"transaction_id" json:"transaction_id,omitempty"`
	Amount        float64    `db:"amount" json:"amount"`
	Status        string     `db:"status" json:"status"`
	PaidAt        *time.Time `db:"paid_at" json:"paid_at,omitempty"`
	CreatedAt     time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updated_at"`
}