package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Adhityaekp/cinema-ticket/internal/dto"
	"github.com/Adhityaekp/cinema-ticket/internal/repository"
	"github.com/Adhityaekp/cinema-ticket/internal/service"
	"github.com/Adhityaekp/cinema-ticket/pkg/response"
	"github.com/gorilla/mux"
)

type StudioHandler struct {
	studioService *service.StudioService
}

func NewStudioHandler(
	studioService *service.StudioService,
) *StudioHandler {
	return &StudioHandler{
		studioService: studioService,
	}
}

// Create godoc
// @Summary Create studio
// @Description Membuat studio baru
// @Tags Studios
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body dto.CreateStudioRequest true "Create studio request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Router /studios [post]
func (h *StudioHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.CreateStudioRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"bad_request",
			"request tidak valid",
		)
		return
	}

	studio, err := h.studioService.Create(
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
		"studio dan seat berhasil dibuat",
		studio,
	)
}

// FindAll godoc
// @Summary Get all studios
// @Description Mendapatkan semua studio
// @Tags Studios
// @Produce json
// @Success 200 {object} response.Response
// @Router /studios [get]
func (h *StudioHandler) FindAll(
	w http.ResponseWriter,
	r *http.Request,
) {
	studios, err := h.studioService.FindAll(r.Context())
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"internal_server_error",
			"gagal mengambil data studio",
		)
		return
	}

	response.Success(
		w,
		"data studio berhasil diambil",
		studios,
	)
}

// FindByID godoc
// @Summary Get studio by ID
// @Description Mendapatkan studio berdasarkan ID
// @Tags Studios
// @Produce json
// @Param id path string true "Studio ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /studios/{id} [get]
func (h *StudioHandler) FindByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := mux.Vars(r)["id"]

	studio, err := h.studioService.FindByID(
		r.Context(),
		id,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(
				w,
				http.StatusNotFound,
				"not_found",
				"studio tidak ditemukan",
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
		"studio berhasil diambil",
		studio,
	)
}

// Update godoc
// @Summary Update studio
// @Description Mengubah data studio
// @Tags Studios
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param id path string true "Studio ID"
// @Param request body dto.UpdateStudioRequest true "Update studio request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /studios/{id} [put]
func (h *StudioHandler) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := mux.Vars(r)["id"]

	var req dto.UpdateStudioRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"bad_request",
			"request tidak valid",
		)
		return
	}

	studio, err := h.studioService.Update(
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
				"studio tidak ditemukan",
			)
			return
		}

		if errors.Is(
			err,
			repository.ErrStudioHasTransaction,
		) {
			response.Error(
				w,
				http.StatusConflict,
				"conflict",
				"studio sudah digunakan untuk showtime",
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
		"studio berhasil diperbarui",
		studio,
	)
}

// Delete godoc
// @Summary Delete studio
// @Description Menghapus studio
// @Tags Studios
// @Produce json
// @Security CookieAuth
// @Param id path string true "Studio ID"
// @Success 200 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /studios/{id} [delete]
func (h *StudioHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := mux.Vars(r)["id"]

	err := h.studioService.Delete(
		r.Context(),
		id,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(
				w,
				http.StatusNotFound,
				"not_found",
				"studio tidak ditemukan",
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
		"studio berhasil dihapus",
		nil,
	)
}
