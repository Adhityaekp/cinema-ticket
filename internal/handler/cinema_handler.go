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

type CinemaHandler struct {
	service *service.CinemaService
}

func NewCinemaHandler(service *service.CinemaService) *CinemaHandler {
	return &CinemaHandler{
		service: service,
	}
}

// CreateCinema godoc
// @Summary Create cinema
// @Description Membuat data cinema baru
// @Tags Cinemas
// @Accept json
// @Produce json
// @Param request body dto.CreateCinemaRequest true "Cinema"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /cinemas [post]
func (h *CinemaHandler) Create(w http.ResponseWriter, r *http.Request) {

	var req dto.CreateCinemaRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"request tidak valid",
			err.Error(),
		)
		return
	}

	cinema, err := h.service.Create(r.Context(), req)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"gagal membuat cinema",
			err.Error(),
		)
		return
	}

	response.Created(
		w,
		"cinema berhasil dibuat",
		cinema,
	)
}

// GetAllCinema godoc
// @Summary Get all cinemas
// @Description Mengambil seluruh data cinema
// @Tags Cinemas
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /cinemas [get]
func (h *CinemaHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	cinemas, err := h.service.FindAll(r.Context())
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"gagal mengambil data cinema",
			err.Error(),
		)
		return
	}

	response.Success(
		w,
		"data cinema berhasil diambil",
		cinemas,
	)
}

// GetCinemaByID godoc
// @Summary Get cinema by ID
// @Description Mengambil detail cinema berdasarkan ID
// @Tags Cinemas
// @Produce json
// @Param id path string true "Cinema ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /cinemas/{id} [get]
func (h *CinemaHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	cinema, err := h.service.FindByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(
				w,
				http.StatusNotFound,
				"cinema tidak ditemukan",
				nil,
			)
			return
		}

		response.Error(
			w,
			http.StatusBadRequest,
			"gagal mengambil cinema",
			err.Error(),
		)
		return
	}

	response.Success(
		w,
		"data cinema berhasil diambil",
		cinema,
	)
}

// UpdateCinema godoc
// @Summary Update cinema
// @Description Mengubah data cinema
// @Tags Cinemas
// @Accept json
// @Produce json
// @Param id path string true "Cinema ID"
// @Param request body dto.UpdateCinemaRequest true "Cinema"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /cinemas/{id} [put]
func (h *CinemaHandler) Update(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	var req dto.UpdateCinemaRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"request tidak valid",
			err.Error(),
		)
		return
	}

	cinema, err := h.service.Update(
		r.Context(),
		id,
		req,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(
				w,
				http.StatusNotFound,
				"cinema tidak ditemukan",
				nil,
			)
			return
		}

		response.Error(
			w,
			http.StatusBadRequest,
			"gagal mengubah cinema",
			err.Error(),
		)
		return
	}

	response.Success(
		w,
		"cinema berhasil diubah",
		cinema,
	)
}

// DeleteCinema godoc
// @Summary Delete cinema
// @Description Menghapus data cinema
// @Tags Cinemas
// @Produce json
// @Param id path string true "Cinema ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 409 {object} response.Response
// @Router /cinemas/{id} [delete]
func (h *CinemaHandler) Delete(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	err := h.service.Delete(r.Context(), id)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			response.Error(
				w,
				http.StatusNotFound,
				"cinema tidak ditemukan",
				nil,
			)
			return
		}

		response.Error(
			w,
			http.StatusConflict,
			"cinema tidak dapat dihapus karena masih digunakan",
			err.Error(),
		)
		return
	}

	response.Success(
		w,
		"cinema berhasil dihapus",
		nil,
	)
}
