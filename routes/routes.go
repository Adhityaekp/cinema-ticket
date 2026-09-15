package routes

import (
	"net/http"

	"github.com/Adhityaekp/cinema-ticket/config"
	"github.com/Adhityaekp/cinema-ticket/internal/handler"
	"github.com/Adhityaekp/cinema-ticket/internal/middleware"
	"github.com/gorilla/mux"
)

func SetupRoutes(
	cfg *config.Config,
	authHandler *handler.AuthHandler,
	showtimeHandler *handler.ShowtimeHandler,
	cinemaHandler *handler.CinemaHandler,
	studioHandler *handler.StudioHandler,
	seatHandler *handler.SeatHandler,
	movieHandler *handler.MovieHandler,
) *mux.Router {

	router := mux.NewRouter()

	// =====================================================
	// AUTH - PUBLIC
	// =====================================================

	router.HandleFunc(
		"/api/auth/register",
		authHandler.Register,
	).Methods("POST")

	router.HandleFunc(
		"/api/auth/verify-email",
		authHandler.VerifyEmail,
	).Methods("GET")

	router.HandleFunc(
		"/api/auth/login",
		authHandler.Login,
	).Methods("POST")

	// =====================================================
	// CINEMAS
	// =====================================================

	// CUSTOMER + ADMIN
	router.HandleFunc(
		"/api/cinemas",
		cinemaHandler.GetAll,
	).Methods("GET")

	router.HandleFunc(
		"/api/cinemas/{id}",
		cinemaHandler.GetByID,
	).Methods("GET")

	// ADMIN ONLY
	router.Handle(
		"/api/cinemas",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(cinemaHandler.Create),
			),
		),
	).Methods("POST")

	router.Handle(
		"/api/cinemas/{id}",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(cinemaHandler.Update),
			),
		),
	).Methods("PUT")

	router.Handle(
		"/api/cinemas/{id}",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(cinemaHandler.Delete),
			),
		),
	).Methods("DELETE")

	// =====================================================
	// STUDIOS
	// =====================================================
	// Studio GET public
	router.HandleFunc(
		"/api/studios",
		studioHandler.FindAll,
	).Methods("GET")

	router.HandleFunc(
		"/api/studios/{id}",
		studioHandler.FindByID,
	).Methods("GET")

	// Studio admin
	router.Handle(
		"/api/studios",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(studioHandler.Create),
			),
		),
	).Methods("POST")

	router.Handle(
		"/api/studios/{id}",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(studioHandler.Update),
			),
		),
	).Methods("PUT")

	router.Handle(
		"/api/studios/{id}",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(studioHandler.Delete),
			),
		),
	).Methods("DELETE")

	// =====================================================
	// SEATS
	// =====================================================
	// Seat GET public
	router.HandleFunc(
		"/api/seats",
		seatHandler.FindAll,
	).Methods("GET")

	router.HandleFunc(
		"/api/seats/{id}",
		seatHandler.FindByID,
	).Methods("GET")

	// Seat admin
	router.Handle(
		"/api/seats",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(seatHandler.Create),
			),
		),
	).Methods("POST")

	router.Handle(
		"/api/seats/{id}",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(seatHandler.Update),
			),
		),
	).Methods("PUT")

	router.Handle(
		"/api/seats/{id}",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(seatHandler.Delete),
			),
		),
	).Methods("DELETE")

	// =====================================================
	// MOVIES
	// =====================================================
	// Movie GET public
	router.HandleFunc(
		"/api/movies",
		movieHandler.FindAll,
	).Methods("GET")

	router.HandleFunc(
		"/api/movies/{id}",
		movieHandler.FindByID,
	).Methods("GET")

	// Movie CREATE - ADMIN
	router.Handle(
		"/api/movies",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(movieHandler.Create),
			),
		),
	).Methods("POST")

	// Movie UPDATE - ADMIN
	router.Handle(
		"/api/movies/{id}",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(movieHandler.Update),
			),
		),
	).Methods("PUT")

	// Movie DELETE - ADMIN
	router.Handle(
		"/api/movies/{id}",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(movieHandler.Delete),
			),
		),
	).Methods("DELETE")

	// =====================================================
	// SHOWTIMES
	// =====================================================

	// CUSTOMER + ADMIN
	router.HandleFunc(
		"/api/showtimes",
		showtimeHandler.FindAll,
	).Methods("GET")

	router.HandleFunc(
		"/api/showtimes/{id}",
		showtimeHandler.FindByID,
	).Methods("GET")

	// ADMIN ONLY
	router.Handle(
		"/api/showtimes",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(showtimeHandler.Create),
			),
		),
	).Methods("POST")

	router.Handle(
		"/api/showtimes/{id}",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(showtimeHandler.Update),
			),
		),
	).Methods("PUT")

	router.Handle(
		"/api/showtimes/{id}",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			middleware.RoleMiddleware("ADMIN")(
				http.HandlerFunc(showtimeHandler.Delete),
			),
		),
	).Methods("DELETE")

	return router
}
