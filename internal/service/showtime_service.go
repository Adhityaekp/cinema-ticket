package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/Adhityaekp/cinema-ticket/internal/dto"
	"github.com/Adhityaekp/cinema-ticket/internal/model"
	"github.com/Adhityaekp/cinema-ticket/internal/repository"
	"github.com/google/uuid"
)

type ShowtimeService struct {
	showtimeRepo *repository.ShowtimeRepository
}

func NewShowtimeService(
	showtimeRepo *repository.ShowtimeRepository,
) *ShowtimeService {
	return &ShowtimeService{
		showtimeRepo: showtimeRepo,
	}
}

func (s *ShowtimeService) Create(
	ctx context.Context,
	req dto.CreateShowtimeRequest,
) (*model.Showtime, error) {

	if req.EndTime.Before(req.StartTime) ||
		req.EndTime.Equal(req.StartTime) {
		return nil, errors.New("end time harus setelah start time")
	}

	movieID, err := uuid.Parse(req.MovieID)
	if err != nil {
		return nil, errors.New("movie_id tidak valid")
	}

	cinemaID, err := uuid.Parse(req.CinemaID)
	if err != nil {
		return nil, errors.New("cinema_id tidak valid")
	}

	studioID, err := uuid.Parse(req.StudioID)
	if err != nil {
		return nil, errors.New("studio_id tidak valid")
	}

	showtime := &model.Showtime{
		MovieID:   movieID,
		CinemaID:  cinemaID,
		StudioID:  studioID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Price:     req.Price,
		Status:    "SCHEDULED",
	}

	if err := s.showtimeRepo.Create(ctx, showtime); err != nil {
		return nil, err
	}

	return showtime, nil
}

func (s *ShowtimeService) FindAll(
	ctx context.Context,
) ([]model.Showtime, error) {
	return s.showtimeRepo.FindAll(ctx)
}

func (s *ShowtimeService) FindByID(
	ctx context.Context,
	id string,
) (*model.Showtime, error) {

	showtimeID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id jadwal tayang tidak valid")
	}

	return s.showtimeRepo.FindByID(ctx, showtimeID)
}

func (s *ShowtimeService) Update(
	ctx context.Context,
	id string,
	req dto.UpdateShowtimeRequest,
) (*model.Showtime, error) {

	showtimeID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id jadwal tayang tidak valid")
	}

	if req.EndTime.Before(req.StartTime) ||
		req.EndTime.Equal(req.StartTime) {
		return nil, errors.New("end time harus setelah start time")
	}

	movieID, err := uuid.Parse(req.MovieID)
	if err != nil {
		return nil, errors.New("movie_id tidak valid")
	}

	cinemaID, err := uuid.Parse(req.CinemaID)
	if err != nil {
		return nil, errors.New("cinema_id tidak valid")
	}

	studioID, err := uuid.Parse(req.StudioID)
	if err != nil {
		return nil, errors.New("studio_id tidak valid")
	}

	status := strings.ToUpper(req.Status)

	showtime := &model.Showtime{
		ID:        showtimeID,
		MovieID:   movieID,
		CinemaID:  cinemaID,
		StudioID:  studioID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Price:     req.Price,
		Status:    status,
	}

	if err := s.showtimeRepo.Update(ctx, showtime); err != nil {
		return nil, err
	}

	return showtime, nil
}

func (s *ShowtimeService) Delete(
	ctx context.Context,
	id string,
) error {

	showtimeID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("id jadwal tayang tidak valid")
	}

	return s.showtimeRepo.Delete(ctx, showtimeID)
}

// Prevent unused import when this file is expanded later.
var _ = time.Now