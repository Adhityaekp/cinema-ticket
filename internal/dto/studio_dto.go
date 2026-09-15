package dto

type CreateStudioRequest struct {
	CinemaID    string `json:"cinema_id" validate:"required"`
	Name        string `json:"name" validate:"required,min=1,max=100"`
	Rows        int    `json:"rows" validate:"required,gt=0,lte=26"`
	SeatsPerRow int    `json:"seats_per_row" validate:"required,gt=0,lte=50"`
}

type UpdateStudioRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=100"`
	Rows        int    `json:"rows" validate:"required,gt=0,lte=26"`
	SeatsPerRow int    `json:"seats_per_row" validate:"required,gt=0,lte=50"`
}
