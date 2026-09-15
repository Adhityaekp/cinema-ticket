package repository

import (
	"context"
	"database/sql"

	"github.com/Adhityaekp/cinema-ticket/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type MovieRepository struct {
	db *sqlx.DB
}

func NewMovieRepository(db *sqlx.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) Create(
	ctx context.Context,
	movie *model.Movie,
) error {
	query := `
		INSERT INTO movies (
			id,
			title,
			description,
			duration_minutes,
			genre,
			age_rating
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING
			id,
			title,
			description,
			duration_minutes,
			genre,
			age_rating,
			created_at,
			updated_at
	`

	return r.db.GetContext(
		ctx,
		movie,
		query,
		movie.ID,
		movie.Title,
		movie.Description,
		movie.DurationMinutes,
		movie.Genre,
		movie.AgeRating,
	)
}

func (r *MovieRepository) FindAll(
	ctx context.Context,
) ([]model.Movie, error) {
	var movies []model.Movie

	query := `
		SELECT
			id,
			title,
			description,
			duration_minutes,
			genre,
			age_rating,
			created_at,
			updated_at
		FROM movies
		ORDER BY title ASC
	`

	err := r.db.SelectContext(
		ctx,
		&movies,
		query,
	)

	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MovieRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*model.Movie, error) {
	var movie model.Movie

	query := `
		SELECT
			id,
			title,
			description,
			duration_minutes,
			genre,
			age_rating,
			created_at,
			updated_at
		FROM movies
		WHERE id = $1
	`

	err := r.db.GetContext(
		ctx,
		&movie,
		query,
		id,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *MovieRepository) Update(
	ctx context.Context,
	movie *model.Movie,
) error {
	query := `
		UPDATE movies
		SET
			title = $1,
			description = $2,
			duration_minutes = $3,
			genre = $4,
			age_rating = $5,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $6
		RETURNING
			id,
			title,
			description,
			duration_minutes,
			genre,
			age_rating,
			created_at,
			updated_at
	`

	return r.db.GetContext(
		ctx,
		movie,
		query,
		movie.Title,
		movie.Description,
		movie.DurationMinutes,
		movie.Genre,
		movie.AgeRating,
		movie.ID,
	)
}

func (r *MovieRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	query := `
		DELETE FROM movies
		WHERE id = $1
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
