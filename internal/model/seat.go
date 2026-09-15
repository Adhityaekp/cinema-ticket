package model

import (
	"time"

	"github.com/google/uuid"
)

type Seat struct {
	ID         uuid.UUID `db:"id" json:"id"`
	StudioID   uuid.UUID `db:"studio_id" json:"studio_id"`
	SeatNumber string    `db:"seat_number" json:"seat_number"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}
