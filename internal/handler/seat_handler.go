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

type SeatHandler struct {
	seatService *service.SeatService
}

func NewSeatHandler(
	seatService *service.SeatService,
) *SeatHandler {
	return &SeatHandler{
		seatService: seatService,
	}
}

// Create godoc
// @Summary Create seat
// @Description Membuat kursi baru pada studio
// @Tags Seats
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body dto.CreateSeatRequest true "Create seat request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Router /seats [post]
func (h *SeatHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.CreateSeatRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"bad_request",
			"request tidak valid",
		)
		return
	}

	seat, err := h.seatService.Create(
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
		"seat berhasil dibuat",
		seat,
	)
}

// FindAll godoc
// @Summary Get all seats
// @Description Mendapatkan semua kursi
// @Tags Seats
// @Produce json
// @Success 200 {object} response.Response
// @Router /seats [get]
func (h *SeatHandler) FindAll(
	w http.ResponseWriter,
	r *http.Request,
) {
	seats, err := h.seatService.FindAll(
		r.Context(),
	)

	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"internal_server_error",
			"gagal mengambil data seat",
		)
		return
	}

	response.Success(
		w,
		"data seat berhasil diambil",
		seats,
	)
}

// FindByID godoc
// @Summary Get seat by ID
// @Description Mendapatkan seat berdasarkan ID
// @Tags Seats
// @Produce json
// @Param id path string true "Seat ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /seats/{id} [get]
func (h *SeatHandler) FindByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := mux.Vars(r)["id"]

	seat, err := h.seatService.FindByID(
		r.Context(),
		id,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(
				w,
				http.StatusNotFound,
				"not_found",
				"seat tidak ditemukan",
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
		"seat berhasil diambil",
		seat,
	)
}

// FindByStudioID godoc
// @Summary Get seats by studio
// @Description Mendapatkan semua kursi berdasarkan studio
// @Tags Seats
// @Produce json
// @Param studio_id query string true "Studio ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /seats [get]
func (h *SeatHandler) FindByStudioID(
	w http.ResponseWriter,
	r *http.Request,
) {
	studioID := r.URL.Query().Get("studio_id")

	if studioID == "" {
		response.Error(
			w,
			http.StatusBadRequest,
			"bad_request",
			"studio_id wajib diisi",
		)
		return
	}

	seats, err := h.seatService.FindByStudioID(
		r.Context(),
		studioID,
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

	response.Success(
		w,
		"data seat berhasil diambil",
		seats,
	)
}

// Update godoc
// @Summary Update seat
// @Description Mengubah data seat
// @Tags Seats
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param id path string true "Seat ID"
// @Param request body dto.UpdateSeatRequest true "Update seat request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /seats/{id} [put]
func (h *SeatHandler) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := mux.Vars(r)["id"]

	var req dto.UpdateSeatRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"bad_request",
			"request tidak valid",
		)
		return
	}

	seat, err := h.seatService.Update(
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
				"seat tidak ditemukan",
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
		"seat berhasil diperbarui",
		seat,
	)
}

// Delete godoc
// @Summary Delete seat
// @Description Menghapus seat
// @Tags Seats
// @Produce json
// @Security CookieAuth
// @Param id path string true "Seat ID"
// @Success 200 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /seats/{id} [delete]
func (h *SeatHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := mux.Vars(r)["id"]

	err := h.seatService.Delete(
		r.Context(),
		id,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(
				w,
				http.StatusNotFound,
				"not_found",
				"seat tidak ditemukan",
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
		"seat berhasil dihapus",
		nil,
	)
}
