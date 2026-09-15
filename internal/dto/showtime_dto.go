package dto

import "time"

type CreateShowtimeRequest struct {
	MovieID   string    `json:"movie_id" validate:"required"`
	CinemaID  string    `json:"cinema_id" validate:"required"`
	StudioID  string    `json:"studio_id" validate:"required"`
	StartTime time.Time `json:"start_time" validate:"required"`
	EndTime   time.Time `json:"end_time" validate:"required"`
	Price     float64   `json:"price" validate:"required,gt=0"`
}

type UpdateShowtimeRequest struct {
	MovieID   string    `json:"movie_id" validate:"required"`
	CinemaID  string    `json:"cinema_id" validate:"required"`
	StudioID  string    `json:"studio_id" validate:"required"`
	StartTime time.Time `json:"start_time" validate:"required"`
	EndTime   time.Time `json:"end_time" validate:"required"`
	Price     float64   `json:"price" validate:"required,gt=0"`
	Status    string    `json:"status" validate:"required,oneof=SCHEDULED ONGOING COMPLETED CANCELLED"`
}
