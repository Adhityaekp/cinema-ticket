package model

import (
	"time"

	"github.com/google/uuid"
)

type Showtime struct {
	ID        uuid.UUID `db:"id" json:"id"`
	MovieID   uuid.UUID `db:"movie_id" json:"movie_id"`
	CinemaID  uuid.UUID `db:"cinema_id" json:"cinema_id"`
	StudioID  uuid.UUID `db:"studio_id" json:"studio_id"`
	StartTime time.Time `db:"start_time" json:"start_time"`
	EndTime   time.Time `db:"end_time" json:"end_time"`
	Price     float64   `db:"price" json:"price"`
	Status    string    `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}