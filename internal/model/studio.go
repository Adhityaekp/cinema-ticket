package model

import (
	"time"

	"github.com/google/uuid"
)

type Studio struct {
	ID        uuid.UUID `db:"id" json:"id"`
	CinemaID  uuid.UUID `db:"cinema_id" json:"cinema_id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}