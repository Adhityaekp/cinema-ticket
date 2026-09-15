package database

import (
	"fmt"

	"github.com/Adhityaekp/cinema-ticket/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectPostgres(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
