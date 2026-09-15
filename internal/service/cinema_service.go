package service

import (
	"context"
	"errors"
	"strings"

	"github.com/Adhityaekp/cinema-ticket/internal/dto"
	"github.com/Adhityaekp/cinema-ticket/internal/model"
	"github.com/Adhityaekp/cinema-ticket/internal/repository"
	"github.com/google/uuid"
)

type CinemaService struct {
	repo *repository.CinemaRepository
}

func NewCinemaService(repo *repository.CinemaRepository) *CinemaService {
	return &CinemaService{
		repo: repo,
	}
}

func (s *CinemaService) Create(
	ctx context.Context,
	req dto.CreateCinemaRequest,
) (*model.Cinema, error) {

	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("nama cinema wajib diisi")
	}

	if strings.TrimSpace(req.City) == "" {
		return nil, errors.New("kota wajib diisi")
	}

	if strings.TrimSpace(req.Address) == "" {
		return nil, errors.New("alamat wajib diisi")
	}

	cinema := &model.Cinema{
		ID:      uuid.New(),
		Name:    strings.TrimSpace(req.Name),
		City:    strings.TrimSpace(req.City),
		Address: strings.TrimSpace(req.Address),
	}

	if err := s.repo.Create(ctx, cinema); err != nil {
		return nil, err
	}

	return cinema, nil
}

func (s *CinemaService) FindAll(
	ctx context.Context,
) ([]model.Cinema, error) {
	return s.repo.FindAll(ctx)
}

func (s *CinemaService) FindByID(
	ctx context.Context,
	id string,
) (*model.Cinema, error) {

	cinemaID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id cinema tidak valid")
	}

	return s.repo.FindByID(ctx, cinemaID)
}

func (s *CinemaService) Update(
	ctx context.Context,
	id string,
	req dto.UpdateCinemaRequest,
) (*model.Cinema, error) {

	cinemaID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id cinema tidak valid")
	}

	cinema := &model.Cinema{
		ID:      cinemaID,
		Name:    strings.TrimSpace(req.Name),
		City:    strings.TrimSpace(req.City),
		Address: strings.TrimSpace(req.Address),
	}

	if cinema.Name == "" {
		return nil, errors.New("nama cinema wajib diisi")
	}

	if cinema.City == "" {
		return nil, errors.New("kota wajib diisi")
	}

	if cinema.Address == "" {
		return nil, errors.New("alamat wajib diisi")
	}

	if err := s.repo.Update(ctx, cinema); err != nil {
		return nil, err
	}

	return cinema, nil
}

func (s *CinemaService) Delete(
	ctx context.Context,
	id string,
) error {

	cinemaID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("id cinema tidak valid")
	}

	return s.repo.Delete(ctx, cinemaID)
}
