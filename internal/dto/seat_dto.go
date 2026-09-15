package dto

type CreateSeatRequest struct {
	StudioID   string `json:"studio_id" validate:"required"`
	SeatNumber string `json:"seat_number" validate:"required,min=1,max=20"`
}

type UpdateSeatRequest struct {
	StudioID   string `json:"studio_id" validate:"required"`
	SeatNumber string `json:"seat_number" validate:"required,min=1,max=20"`
}