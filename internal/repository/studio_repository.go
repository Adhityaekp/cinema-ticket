package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/Adhityaekp/cinema-ticket/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var ErrStudioHasTransaction = errors.New(
	"studio sudah digunakan untuk transaksi",
)

type StudioRepository struct {
	db *sqlx.DB
}

func NewStudioRepository(db *sqlx.DB) *StudioRepository {
	return &StudioRepository{
		db: db,
	}
}

func insertSeats(
	ctx context.Context,
	tx *sqlx.Tx,
	studioID uuid.UUID,
	rows int,
	seatsPerRow int,
) error {
	query := `
		INSERT INTO seats (
			id,
			studio_id,
			seat_number
		)
		VALUES ($1, $2, $3)
	`

	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for row := 0; row < rows; row++ {
		rowLetter := string(rune('A' + row))

		for number := 1; number <= seatsPerRow; number++ {
			seatNumber := rowLetter + strconv.Itoa(number)

			_, err := stmt.ExecContext(
				ctx,
				uuid.New(),
				studioID,
				seatNumber,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *StudioRepository) Create(
	ctx context.Context,
	studio *model.Studio,
	rows int,
	seatsPerRow int,
) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	studioQuery := `
		INSERT INTO studios (
			id,
			cinema_id,
			name
		)
		VALUES ($1, $2, $3)
		RETURNING
			id,
			cinema_id,
			name,
			created_at,
			updated_at
	`

	if err = tx.GetContext(
		ctx,
		studio,
		studioQuery,
		studio.ID,
		studio.CinemaID,
		studio.Name,
	); err != nil {
		return err
	}

	if err = insertSeats(
		ctx,
		tx,
		studio.ID,
		rows,
		seatsPerRow,
	); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *StudioRepository) FindAll(
	ctx context.Context,
) ([]model.Studio, error) {
	var studios []model.Studio

	query := `
		SELECT
			id,
			cinema_id,
			name,
			created_at,
			updated_at
		FROM studios
		ORDER BY name ASC
	`

	err := r.db.SelectContext(
		ctx,
		&studios,
		query,
	)

	if err != nil {
		return nil, err
	}

	return studios, nil
}

func (r *StudioRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*model.Studio, error) {
	var studio model.Studio

	query := `
		SELECT
			id,
			cinema_id,
			name,
			created_at,
			updated_at
		FROM studios
		WHERE id = $1
	`

	err := r.db.GetContext(
		ctx,
		&studio,
		query,
		id,
	)

	if err != nil {
		return nil, err
	}

	return &studio, nil
}

func (r *StudioRepository) FindSeats(
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

func (r *StudioRepository) HasShowtimes(
	ctx context.Context,
	studioID uuid.UUID,
) (bool, error) {
	var exists bool

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM showtimes
			WHERE studio_id = $1
		)
	`

	err := r.db.GetContext(
		ctx,
		&exists,
		query,
		studioID,
	)

	return exists, err
}

func (r *StudioRepository) Update(
	ctx context.Context,
	studio *model.Studio,
	rows int,
	seatsPerRow int,
) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	var hasShowtimes bool

	queryCheck := `
		SELECT EXISTS (
			SELECT 1
			FROM showtimes
			WHERE studio_id = $1
		)
	`

	err = tx.GetContext(
		ctx,
		&hasShowtimes,
		queryCheck,
		studio.ID,
	)

	if err != nil {
		return err
	}

	if hasShowtimes {
		return ErrStudioHasTransaction
	}

	updateQuery := `
		UPDATE studios
		SET
			name = $1,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $2
		RETURNING
			id,
			cinema_id,
			name,
			created_at,
			updated_at
	`

	err = tx.GetContext(
		ctx,
		studio,
		updateQuery,
		studio.Name,
		studio.ID,
	)

	if err != nil {
		return err
	}

	// Delete seats lama
	_, err = tx.ExecContext(
		ctx,
		`DELETE FROM seats WHERE studio_id = $1`,
		studio.ID,
	)

	if err != nil {
		return err
	}

	// Generate seats baru
	if err = insertSeats(
		ctx,
		tx,
		studio.ID,
		rows,
		seatsPerRow,
	); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *StudioRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	var hasShowtimes bool

	err = tx.GetContext(
		ctx,
		&hasShowtimes,
		`
			SELECT EXISTS (
				SELECT 1
				FROM showtimes
				WHERE studio_id = $1
			)
		`,
		id,
	)

	if err != nil {
		return err
	}

	if hasShowtimes {
		return ErrStudioHasTransaction
	}

	result, err := tx.ExecContext(
		ctx,
		`DELETE FROM studios WHERE id = $1`,
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

	// Seat akan ikut terhapus jika FK seats -> studios
	// menggunakan ON DELETE CASCADE.

	return tx.Commit()
}
