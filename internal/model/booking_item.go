package model

import (
	"time"

	"github.com/google/uuid"
)

type BookingItem struct {
	ID             uuid.UUID `db:"id" json:"id"`
	BookingID      uuid.UUID `db:"booking_id" json:"booking_id"`
	ShowtimeSeatID uuid.UUID `db:"showtime_seat_id" json:"showtime_seat_id"`
	Price          float64   `db:"price" json:"price"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}