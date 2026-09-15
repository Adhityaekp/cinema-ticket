package dto

type CreateMovieRequest struct {
	Title           string  `json:"title" validate:"required,min=1,max=200"`
	Description     *string `json:"description"`
	DurationMinutes int     `json:"duration_minutes" validate:"required,gt=0"`
	Genre           *string `json:"genre"`
	AgeRating       *string `json:"age_rating"`
}

type UpdateMovieRequest struct {
	Title           string  `json:"title" validate:"required,min=1,max=200"`
	Description     *string `json:"description"`
	DurationMinutes int     `json:"duration_minutes" validate:"required,gt=0"`
	Genre           *string `json:"genre"`
	AgeRating       *string `json:"age_rating"`
}
