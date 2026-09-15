package model

import (
	"time"

	"github.com/google/uuid"
)

type ShowtimeSeat struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	ShowtimeID uuid.UUID  `db:"showtime_id" json:"showtime_id"`
	SeatID     uuid.UUID  `db:"seat_id" json:"seat_id"`
	Status     string     `db:"status" json:"status"`
	HeldUntil  *time.Time `db:"held_until" json:"held_until,omitempty"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updated_at"`
}