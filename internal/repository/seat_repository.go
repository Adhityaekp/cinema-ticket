package repository

import (
	"context"
	"database/sql"

	"github.com/Adhityaekp/cinema-ticket/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SeatRepository struct {
	db *sqlx.DB
}

func NewSeatRepository(db *sqlx.DB) *SeatRepository {
	return &SeatRepository{
		db: db,
	}
}

func (r *SeatRepository) Create(
	ctx context.Context,
	seat *model.Seat,
) error {
	query := `
		INSERT INTO seats (
			id,
			studio_id,
			seat_number
		)
		VALUES ($1, $2, $3)
		RETURNING
			id,
			studio_id,
			seat_number,
			created_at,
			updated_at
	`

	return r.db.GetContext(
		ctx,
		seat,
		query,
		seat.ID,
		seat.StudioID,
		seat.SeatNumber,
	)
}

func (r *SeatRepository) FindAll(
	ctx context.Context,
) ([]model.Seat, error) {
	var seats []model.Seat

	query := `
		SELECT
			id,
			studio_id,
			seat_number,
			created_at,
			updated_at
		FROM seats
		ORDER BY studio_id, seat_number ASC
	`

	err := r.db.SelectContext(ctx, &seats, query)
	if err != nil {
		return nil, err
	}

	return seats, nil
}

func (r *SeatRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*model.Seat, error) {
	var seat model.Seat

	query := `
		SELECT
			id,
			studio_id,
			seat_number,
			created_at,
			updated_at
		FROM seats
		WHERE id = $1
	`

	err := r.db.GetContext(ctx, &seat, query, id)
	if err != nil {
		return nil, err
	}

	return &seat, nil
}

func (r *SeatRepository) FindByStudioID(
	ctx context.Context,
	studioID uuid.UUID,
) ([]model.Seat, error) {
	var seats []model.Seat

	query := `
		SELECT
			id,
			studio_id,
			seat_number,
			created_at,
			updated_at
		FROM seats
		WHERE studio_id = $1
		ORDER BY seat_number ASC
	`

	err := r.db.SelectContext(
		ctx,
		&seats,
		query,
		studioID,
	)

	if err != nil {
		return nil, err
	}

	return seats, nil
}

func (r *SeatRepository) Update(
	ctx context.Context,
	seat *model.Seat,
) error {
	query := `
		UPDATE seats
		SET
			studio_id = $1,
			seat_number = $2,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $3
		RETURNING
			id,
			studio_id,
			seat_number,
			created_at,
			updated_at
	`

	return r.db.GetContext(
		ctx,
		seat,
		query,
		seat.StudioID,
		seat.SeatNumber,
		seat.ID,
	)
}

func (r *SeatRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	query := `
		DELETE FROM seats
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