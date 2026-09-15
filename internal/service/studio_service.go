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

type StudioService struct {
	studioRepo *repository.StudioRepository
}

func NewStudioService(
	studioRepo *repository.StudioRepository,
) *StudioService {
	return &StudioService{
		studioRepo: studioRepo,
	}
}

func (s *StudioService) Create(
	ctx context.Context,
	req dto.CreateStudioRequest,
) (*model.Studio, error) {
	cinemaID, err := uuid.Parse(req.CinemaID)
	if err != nil {
		return nil, errors.New("cinema_id tidak valid")
	}

	name := strings.TrimSpace(req.Name)

	if name == "" {
		return nil, errors.New("nama studio wajib diisi")
	}

	if req.Rows <= 0 || req.Rows > 26 {
		return nil, errors.New("jumlah baris harus antara 1 sampai 26")
	}

	if req.SeatsPerRow <= 0 || req.SeatsPerRow > 50 {
		return nil, errors.New(
			"jumlah kursi per baris harus antara 1 sampai 50",
		)
	}

	studio := &model.Studio{
		ID:       uuid.New(),
		CinemaID: cinemaID,
		Name:     name,
	}

	err = s.studioRepo.Create(
		ctx,
		studio,
		req.Rows,
		req.SeatsPerRow,
	)

	if err != nil {
		return nil, err
	}

	return studio, nil
}

func (s *StudioService) FindAll(
	ctx context.Context,
) ([]model.Studio, error) {
	return s.studioRepo.FindAll(ctx)
}

func (s *StudioService) FindByID(
	ctx context.Context,
	id string,
) (*model.Studio, error) {

	studioID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id studio tidak valid")
	}

	studio, err := s.studioRepo.FindByID(ctx, studioID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}

		return nil, err
	}

	return studio, nil
}

func (s *StudioService) Update(
	ctx context.Context,
	id string,
	req dto.UpdateStudioRequest,
) (*model.Studio, error) {
	studioID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id studio tidak valid")
	}

	name := strings.TrimSpace(req.Name)

	if name == "" {
		return nil, errors.New("nama studio wajib diisi")
	}

	if req.Rows <= 0 || req.Rows > 26 {
		return nil, errors.New("jumlah baris harus antara 1 sampai 26")
	}

	if req.SeatsPerRow <= 0 || req.SeatsPerRow > 50 {
		return nil, errors.New(
			"jumlah kursi per baris harus antara 1 sampai 50",
		)
	}

	studio, err := s.studioRepo.FindByID(
		ctx,
		studioID,
	)

	if err != nil {
		return nil, err
	}

	studio.Name = name

	err = s.studioRepo.Update(
		ctx,
		studio,
		req.Rows,
		req.SeatsPerRow,
	)

	if err != nil {
		return nil, err
	}

	return studio, nil
}

func (s *StudioService) Delete(
	ctx context.Context,
	id string,
) error {

	studioID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("id studio tidak valid")
	}

	return s.studioRepo.Delete(ctx, studioID)
}
