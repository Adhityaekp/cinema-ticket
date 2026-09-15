package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Adhityaekp/cinema-ticket/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ShowtimeRepository struct {
	db *sqlx.DB
}

func NewShowtimeRepository(db *sqlx.DB) *ShowtimeRepository {
	return &ShowtimeRepository{
		db: db,
	}
}

func (r *ShowtimeRepository) Create(
	ctx context.Context,
	showtime *model.Showtime,
) error {
	query := `
		INSERT INTO showtimes (
			movie_id,
			cinema_id,
			studio_id,
			start_time,
			end_time,
			price,
			status
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowxContext(
		ctx,
		query,
		showtime.MovieID,
		showtime.CinemaID,
		showtime.StudioID,
		showtime.StartTime,
		showtime.EndTime,
		showtime.Price,
		showtime.Status,
	).Scan(
		&showtime.ID,
		&showtime.CreatedAt,
		&showtime.UpdatedAt,
	)
}

func (r *ShowtimeRepository) FindAll(
	ctx context.Context,
) ([]model.Showtime, error) {
	var showtimes []model.Showtime

	query := `
		SELECT
			id,
			movie_id,
			cinema_id,
			studio_id,
			start_time,
			end_time,
			price,
			status,
			created_at,
			updated_at
		FROM showtimes
		ORDER BY start_time ASC
	`

	err := r.db.SelectContext(ctx, &showtimes, query)
	if err != nil {
		return nil, err
	}

	return showtimes, nil
}

func (r *ShowtimeRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*model.Showtime, error) {
	var showtime model.Showtime

	query := `
		SELECT
			id,
			movie_id,
			cinema_id,
			studio_id,
			start_time,
			end_time,
			price,
			status,
			created_at,
			updated_at
		FROM showtimes
		WHERE id = $1
	`

	err := r.db.GetContext(ctx, &showtime, query, id)
	if err != nil {
		return nil, err
	}

	return &showtime, nil
}

func (r *ShowtimeRepository) Update(
	ctx context.Context,
	showtime *model.Showtime,
) error {
	query := `
		UPDATE showtimes
		SET
			movie_id = $1,
			cinema_id = $2,
			studio_id = $3,
			start_time = $4,
			end_time = $5,
			price = $6,
			status = $7,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $8
		RETURNING updated_at
	`

	return r.db.QueryRowxContext(
		ctx,
		query,
		showtime.MovieID,
		showtime.CinemaID,
		showtime.StudioID,
		showtime.StartTime,
		showtime.EndTime,
		showtime.Price,
		showtime.Status,
		showtime.ID,
	).Scan(&showtime.UpdatedAt)
}

func (r *ShowtimeRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	query := `
		DELETE FROM showtimes
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
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

func IsNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
