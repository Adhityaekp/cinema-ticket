package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Adhityaekp/cinema-ticket/config"
	"github.com/Adhityaekp/cinema-ticket/internal/handler"
	"github.com/Adhityaekp/cinema-ticket/internal/repository"
	"github.com/Adhityaekp/cinema-ticket/internal/service"
	"github.com/Adhityaekp/cinema-ticket/pkg/database"
	redisClient "github.com/Adhityaekp/cinema-ticket/pkg/redis"
	"github.com/Adhityaekp/cinema-ticket/routes"

	_ "github.com/Adhityaekp/cinema-ticket/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Cinema Ticket API
// @version 1.0
// @description REST API untuk sistem pemesanan tiket bioskop.
// @host localhost:8080
// @BasePath /api
func main() {
	cfg := config.Load()

	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal("Failed to connect PostgreSQL:", err)
	}
	defer db.Close()

	redisClient, err := redisClient.ConnectRedis(cfg)
	if err != nil {
		log.Fatal("Failed to connect Redis:", err)
	}
	defer redisClient.Close()

	log.Println("PostgreSQL connected successfully")
	log.Println("Redis connected successfully")

	// =========================
	// Seed Dummy Users
	// =========================

	ctx := context.Background()

	userRepository := repository.NewUserRepository(db)

	if err := service.SeedDummyUsers(ctx, userRepository); err != nil {
		log.Fatal("Failed to seed dummy users:", err)
	}

	// =========================
	// Repositories, Services, and Handlers
	// =========================

	showtimeRepository := repository.NewShowtimeRepository(db)
	cinemaRepository := repository.NewCinemaRepository(db)
	studioRepository := repository.NewStudioRepository(db)
	seatRepository := repository.NewSeatRepository(db)
	movieRepository := repository.NewMovieRepository(db)

	emailService := service.NewEmailService(
		cfg.SMTPHost,
		cfg.SMTPPort,
		cfg.SMTPUsername,
		cfg.SMTPPassword,
		cfg.SMTPFrom,
		cfg.AppBaseURL,
	)
	showtimeService := service.NewShowtimeService(
		showtimeRepository,
	)
	cinemaService := service.NewCinemaService(cinemaRepository)
	studioService := service.NewStudioService(studioRepository)
	authService := service.NewAuthService(
		userRepository,
		emailService,
		cfg,
	)
	seatService := service.NewSeatService(seatRepository)
	movieService := service.NewMovieService(movieRepository)

	authHandler := handler.NewAuthHandler(authService, cfg)
	showtimeHandler := handler.NewShowtimeHandler(
		showtimeService,
	)
	cinemaHandler := handler.NewCinemaHandler(cinemaService)
	studioHandler := handler.NewStudioHandler(studioService)
	seatHandler := handler.NewSeatHandler(seatService)
	movieHandler := handler.NewMovieHandler(movieService)

	// =========================
	// Routes
	// =========================

	router := routes.SetupRoutes(cfg, authHandler, showtimeHandler, cinemaHandler, studioHandler, seatHandler, movieHandler)

	// Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// =========================
	// Server
	// =========================

	log.Printf("Server running on :%s", cfg.AppPort)

	if err := http.ListenAndServe(":"+cfg.AppPort, router); err != nil {
		log.Fatal(err)
	}
}
