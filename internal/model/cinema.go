package model

import (
	"time"

	"github.com/google/uuid"
)

type Cinema struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	City      string    `db:"city" json:"city"`
	Address   string    `db:"address" json:"address"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
