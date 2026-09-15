package repository

import (
	"context"
	"database/sql"

	"github.com/Adhityaekp/cinema-ticket/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (
			name,
			email,
			password,
			role,
			is_active,
			email_verified_at,
			email_verification_token,
			email_verification_expires_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowxContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		user.IsActive,
		user.EmailVerifiedAt,
		user.EmailVerificationToken,
		user.EmailVerificationExpiresAt,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}

func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*model.User, error) {

	var user model.User

	query := `
		SELECT
			id,
			name,
			email,
			password,
			role,
			is_active,
			email_verified_at,
			email_verification_token,
			email_verification_expires_at,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`

	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByVerificationToken(
	ctx context.Context,
	token string,
) (*model.User, error) {

	var user model.User

	query := `
		SELECT
			id,
			name,
			email,
			password,
			role,
			is_active,
			email_verified_at,
			email_verification_token,
			email_verification_expires_at,
			created_at,
			updated_at
		FROM users
		WHERE email_verification_token = $1
	`

	err := r.db.GetContext(ctx, &user, query, token)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) VerifyEmail(
	ctx context.Context,
	userID uuid.UUID,
) error {

	query := `
		UPDATE users
		SET
			email_verified_at = CURRENT_TIMESTAMP,
			email_verification_token = NULL,
			email_verification_expires_at = NULL,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
