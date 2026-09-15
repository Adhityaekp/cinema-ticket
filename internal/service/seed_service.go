package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Adhityaekp/cinema-ticket/internal/model"
	"github.com/Adhityaekp/cinema-ticket/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

func SeedDummyUsers(ctx context.Context, userRepo *repository.UserRepository) error {
	dummyUsers := []struct {
		Name     string
		Email    string
		Password string
		Role     string
	}{
		{
			Name:     "Dummy Admin",
			Email:    "admin@cinematicket.com",
			Password: "	!",
			Role:     "ADMIN",
		},
		{
			Name:     "Dummy User",
			Email:    "user@cinematicket.com",
			Password: "User123!",
			Role:     "CUSTOMER",
		},
	}

	for _, dummy := range dummyUsers {
		_, err := userRepo.FindByEmail(ctx, dummy.Email)

		if err == nil {
			log.Printf("Dummy user already exists: %s", dummy.Email)
			continue
		}

		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(dummy.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			return err
		}

		now := time.Now()

		user := &model.User{
			Name:            dummy.Name,
			Email:           dummy.Email,
			Password:        string(hashedPassword),
			Role:            dummy.Role,
			IsActive:        true,
			EmailVerifiedAt: &now,
		}

		// Dummy account langsung dianggap sudah verifikasi email.
		// Jadi perlu repository method khusus untuk membuat user verified,
		// atau Create perlu menerima EmailVerifiedAt.
		if err := userRepo.Create(ctx, user); err != nil {
			return err
		}

		log.Printf("Dummy user created: %s", dummy.Email)
	}

	return nil
}
