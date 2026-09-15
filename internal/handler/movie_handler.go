package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Adhityaekp/cinema-ticket/internal/dto"
	"github.com/Adhityaekp/cinema-ticket/internal/service"
	"github.com/Adhityaekp/cinema-ticket/pkg/response"
	"github.com/gorilla/mux"
)

type MovieHandler struct {
	movieService *service.MovieService
}

func NewMovieHandler(
	movieService *service.MovieService,
) *MovieHandler {
	return &MovieHandler{
		movieService: movieService,
	}
}

// Create godoc
// @Summary Create movie
// @Description Membuat data film baru
// @Tags Movies
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body dto.CreateMovieRequest true "Create movie request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Router /movies [post]
func (h *MovieHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.CreateMovieRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"bad_request",
			"request tidak valid",
		)
		return
	}

	movie, err := h.movieService.Create(
		r.Context(),
		req,
	)

	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"bad_request",
			err.Error(),
		)
		return
	}

	response.Created(
		w,
		"movie berhasil dibuat",
		movie,
	)
}

// FindAll godoc
// @Summary Get all movies
// @Description Mendapatkan semua film
// @Tags Movies
// @Produce json
// @Success 200 {object} response.Response
// @Router /movies [get]
func (h *MovieHandler) FindAll(
	w http.ResponseWriter,
	r *http.Request,
) {
	movies, err := h.movieService.FindAll(
		r.Context(),
	)

	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"internal_server_error",
			"gagal mengambil data movie",
		)
		return
	}

	response.Success(
		w,
		"data movie berhasil diambil",
		movies,
	)
}

// FindByID godoc
// @Summary Get movie by ID
// @Description Mendapatkan film berdasarkan ID
// @Tags Movies
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /movies/{id} [get]
func (h *MovieHandler) FindByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := mux.Vars(r)["id"]

	movie, err := h.movieService.FindByID(
		r.Context(),
		id,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(
				w,
				http.StatusNotFound,
				"not_found",
				"movie tidak ditemukan",
			)
			return
		}

		response.Error(
			w,
			http.StatusBadRequest,
			"bad_request",
			err.Error(),
		)
		return
	}

	response.Success(
		w,
		"movie berhasil diambil",
		movie,
	)
}

// Update godoc
// @Summary Update movie
// @Description Mengubah data film
// @Tags Movies
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param id path string true "Movie ID"
// @Param request body dto.UpdateMovieRequest true "Update movie request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /movies/{id} [put]
func (h *MovieHandler) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := mux.Vars(r)["id"]

	var req dto.UpdateMovieRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"bad_request",
			"request tidak valid",
		)
		return
	}

	movie, err := h.movieService.Update(
		r.Context(),
		id,
		req,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(
				w,
				http.StatusNotFound,
				"not_found",
				"movie tidak ditemukan",
			)
			return
		}

		response.Error(
			w,
			http.StatusBadRequest,
			"bad_request",
			err.Error(),
		)
		return
	}

	response.Success(
		w,
		"movie berhasil diperbarui",
		movie,
	)
}

// Delete godoc
// @Summary Delete movie
// @Description Menghapus data film
// @Tags Movies
// @Produce json
// @Security CookieAuth
// @Param id path string true "Movie ID"
// @Success 200 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /movies/{id} [delete]
func (h *MovieHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := mux.Vars(r)["id"]

	err := h.movieService.Delete(
		r.Context(),
		id,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(
				w,
				http.StatusNotFound,
				"not_found",
				"movie tidak ditemukan",
			)
			return
		}

		response.Error(
			w,
			http.StatusBadRequest,
			"bad_request",
			err.Error(),
		)
		return
	}

	response.Success(
		w,
		"movie berhasil dihapus",
		nil,
	)
}
