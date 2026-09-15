package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Adhityaekp/cinema-ticket/internal/dto"
	"github.com/Adhityaekp/cinema-ticket/internal/service"
	"github.com/Adhityaekp/cinema-ticket/pkg/response"
	"github.com/gorilla/mux"
)

type ShowtimeHandler struct {
	showtimeService *service.ShowtimeService
}

func NewShowtimeHandler(
	showtimeService *service.ShowtimeService,
) *ShowtimeHandler {
	return &ShowtimeHandler{
		showtimeService: showtimeService,
	}
}

// Create godoc
// @Summary Create jadwal tayang
// @Description Membuat jadwal tayang baru.
// @Tags Showtimes
// @Accept json
// @Produce json
// @Param request body dto.CreateShowtimeRequest true "Create showtime request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /showtimes [post]
func (h *ShowtimeHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.CreateShowtimeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"Request tidak valid",
			nil,
		)
		return
	}

	showtime, err := h.showtimeService.Create(
		r.Context(),
		req,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
			nil,
		)
		return
	}

	response.Created(
		w,
		"Jadwal tayang berhasil dibuat",
		showtime,
	)
}

// FindAll godoc
// @Summary Get semua jadwal tayang
// @Description Mendapatkan semua jadwal tayang.
// @Tags Showtimes
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /showtimes [get]
func (h *ShowtimeHandler) FindAll(
	w http.ResponseWriter,
	r *http.Request,
) {
	showtimes, err := h.showtimeService.FindAll(
		r.Context(),
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"Gagal mengambil data jadwal tayang",
			nil,
		)
		return
	}

	response.Success(
		w,
		"Data jadwal tayang berhasil diambil",
		showtimes,
	)
}

// FindByID godoc
// @Summary Get jadwal tayang berdasarkan ID
// @Description Mendapatkan detail jadwal tayang berdasarkan ID.
// @Tags Showtimes
// @Produce json
// @Param id path string true "Showtime ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /showtimes/{id} [get]
func (h *ShowtimeHandler) FindByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := mux.Vars(r)["id"]

	showtime, err := h.showtimeService.FindByID(
		r.Context(),
		id,
	)
	if err != nil {
		if errors.Is(err, errors.New("id jadwal tayang tidak valid")) {
			response.Error(
				w,
				http.StatusBadRequest,
				err.Error(),
				nil,
			)
			return
		}

		response.Error(
			w,
			http.StatusNotFound,
			"Jadwal tayang tidak ditemukan",
			nil,
		)
		return
	}

	response.Success(
		w,
		"Data jadwal tayang berhasil diambil",
		showtime,
	)
}

// Update godoc
// @Summary Update jadwal tayang
// @Description Mengubah data jadwal tayang.
// @Tags Showtimes
// @Accept json
// @Produce json
// @Param id path string true "Showtime ID"
// @Param request body dto.UpdateShowtimeRequest true "Update showtime request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /showtimes/{id} [put]
func (h *ShowtimeHandler) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := mux.Vars(r)["id"]

	var req dto.UpdateShowtimeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"Request tidak valid",
			nil,
		)
		return
	}

	showtime, err := h.showtimeService.Update(
		r.Context(),
		id,
		req,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
			nil,
		)
		return
	}

	response.Success(
		w,
		"Jadwal tayang berhasil diperbarui",
		showtime,
	)
}

// Delete godoc
// @Summary Delete jadwal tayang
// @Description Menghapus jadwal tayang.
// @Tags Showtimes
// @Produce json
// @Param id path string true "Showtime ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /showtimes/{id} [delete]
func (h *ShowtimeHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := mux.Vars(r)["id"]

	err := h.showtimeService.Delete(
		r.Context(),
		id,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
			nil,
		)
		return
	}

	response.Success(
		w,
		"Jadwal tayang berhasil dihapus",
		nil,
	)
}
