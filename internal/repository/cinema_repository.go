package repository

import (
	"context"
	"database/sql"

	"github.com/Adhityaekp/cinema-ticket/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CinemaRepository struct {
	db *sqlx.DB
}

func NewCinemaRepository(db *sqlx.DB) *CinemaRepository {
	return &CinemaRepository{db: db}
}

func (r *CinemaRepository) Create(
	ctx context.Context,
	cinema *model.Cinema,
) error {
	query := `
		INSERT INTO cinemas (
			id,
			name,
			city,
			address
		)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id,
			name,
			city,
			address,
			created_at,
			updated_at
	`

	return r.db.GetContext(
		ctx,
		cinema,
		query,
		cinema.ID,
		cinema.Name,
		cinema.City,
		cinema.Address,
	)
}

func (r *CinemaRepository) FindAll(
	ctx context.Context,
) ([]model.Cinema, error) {
	var cinemas []model.Cinema

	query := `
		SELECT
			id,
			name,
			city,
			address,
			created_at,
			updated_at
		FROM cinemas
		ORDER BY name ASC
	`

	err := r.db.SelectContext(ctx, &cinemas, query)
	return cinemas, err
}

func (r *CinemaRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*model.Cinema, error) {
	var cinema model.Cinema

	query := `
		SELECT
			id,
			name,
			city,
			address,
			created_at,
			updated_at
		FROM cinemas
		WHERE id = $1
	`

	err := r.db.GetContext(ctx, &cinema, query, id)
	if err != nil {
		return nil, err
	}

	return &cinema, nil
}

func (r *CinemaRepository) Update(
	ctx context.Context,
	cinema *model.Cinema,
) error {
	query := `
		UPDATE cinemas
		SET
			name = $1,
			city = $2,
			address = $3,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $4
		RETURNING
			id,
			name,
			city,
			address,
			created_at,
			updated_at
	`

	return r.db.GetContext(
		ctx,
		cinema,
		query,
		cinema.Name,
		cinema.City,
		cinema.Address,
		cinema.ID,
	)
}

func (r *CinemaRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	query := `
		DELETE FROM cinemas
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
