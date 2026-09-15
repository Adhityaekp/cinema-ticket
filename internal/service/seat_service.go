package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/Adhityaekp/cinema-ticket/internal/dto"
	"github.com/Adhityaekp/cinema-ticket/internal/model"
	"github.com/Adhityaekp/cinema-ticket/internal/repository"
	"github.com/google/uuid"
)

type SeatService struct {
	seatRepo *repository.SeatRepository
}

func NewSeatService(
	seatRepo *repository.SeatRepository,
) *SeatService {
	return &SeatService{
		seatRepo: seatRepo,
	}
}

func (s *SeatService) Create(
	ctx context.Context,
	req dto.CreateSeatRequest,
) (*model.Seat, error) {
	studioID, err := uuid.Parse(req.StudioID)
	if err != nil {
		return nil, errors.New("studio_id tidak valid")
	}

	seatNumber := strings.TrimSpace(req.SeatNumber)

	if seatNumber == "" {
		return nil, errors.New("nomor kursi wajib diisi")
	}

	seat := &model.Seat{
		ID:         uuid.New(),
		StudioID:   studioID,
		SeatNumber: seatNumber,
	}

	err = s.seatRepo.Create(ctx, seat)
	if err != nil {
		return nil, err
	}

	return seat, nil
}

func (s *SeatService) FindAll(
	ctx context.Context,
) ([]model.Seat, error) {
	return s.seatRepo.FindAll(ctx)
}

func (s *SeatService) FindByID(
	ctx context.Context,
	id string,
) (*model.Seat, error) {
	seatID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id seat tidak valid")
	}

	seat, err := s.seatRepo.FindByID(ctx, seatID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}

		return nil, err
	}

	return seat, nil
}

func (s *SeatService) FindByStudioID(
	ctx context.Context,
	studioID string,
) ([]model.Seat, error) {
	id, err := uuid.Parse(studioID)
	if err != nil {
		return nil, errors.New("studio_id tidak valid")
	}

	return s.seatRepo.FindByStudioID(ctx, id)
}

func (s *SeatService) Update(
	ctx context.Context,
	id string,
	req dto.UpdateSeatRequest,
) (*model.Seat, error) {
	seatID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id seat tidak valid")
	}

	studioID, err := uuid.Parse(req.StudioID)
	if err != nil {
		return nil, errors.New("studio_id tidak valid")
	}

	seatNumber := strings.TrimSpace(req.SeatNumber)

	if seatNumber == "" {
		return nil, errors.New("nomor kursi wajib diisi")
	}

	seat := &model.Seat{
		ID:         seatID,
		StudioID:   studioID,
		SeatNumber: seatNumber,
	}

	err = s.seatRepo.Update(ctx, seat)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}

		return nil, err
	}

	return seat, nil
}

func (s *SeatService) Delete(
	ctx context.Context,
	id string,
) error {
	seatID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("id seat tidak valid")
	}

	return s.seatRepo.Delete(ctx, seatID)
}
