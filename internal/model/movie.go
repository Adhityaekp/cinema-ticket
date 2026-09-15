package model

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	ID              uuid.UUID `db:"id" json:"id"`
	Title           string    `db:"title" json:"title"`
	Description     *string   `db:"description" json:"description,omitempty"`
	DurationMinutes int       `db:"duration_minutes" json:"duration_minutes"`
	Genre           *string   `db:"genre" json:"genre,omitempty"`
	AgeRating       *string   `db:"age_rating" json:"age_rating,omitempty"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}