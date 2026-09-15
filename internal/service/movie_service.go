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

type MovieService struct {
	movieRepo *repository.MovieRepository
}

func NewMovieService(
	movieRepo *repository.MovieRepository,
) *MovieService {
	return &MovieService{
		movieRepo: movieRepo,
	}
}

func (s *MovieService) Create(
	ctx context.Context,
	req dto.CreateMovieRequest,
) (*model.Movie, error) {
	title := strings.TrimSpace(req.Title)

	if title == "" {
		return nil, errors.New("judul film wajib diisi")
	}

	if req.DurationMinutes <= 0 {
		return nil, errors.New(
			"durasi film harus lebih dari 0 menit",
		)
	}

	movie := &model.Movie{
		ID:              uuid.New(),
		Title:           title,
		Description:     req.Description,
		DurationMinutes: req.DurationMinutes,
		Genre:           req.Genre,
		AgeRating:       req.AgeRating,
	}

	err := s.movieRepo.Create(
		ctx,
		movie,
	)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *MovieService) FindAll(
	ctx context.Context,
) ([]model.Movie, error) {
	return s.movieRepo.FindAll(ctx)
}

func (s *MovieService) FindByID(
	ctx context.Context,
	id string,
) (*model.Movie, error) {
	movieID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id movie tidak valid")
	}

	movie, err := s.movieRepo.FindByID(
		ctx,
		movieID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}

		return nil, err
	}

	return movie, nil
}

func (s *MovieService) Update(
	ctx context.Context,
	id string,
	req dto.UpdateMovieRequest,
) (*model.Movie, error) {
	movieID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id movie tidak valid")
	}

	title := strings.TrimSpace(req.Title)

	if title == "" {
		return nil, errors.New("judul film wajib diisi")
	}

	if req.DurationMinutes <= 0 {
		return nil, errors.New(
			"durasi film harus lebih dari 0 menit",
		)
	}

	movie := &model.Movie{
		ID:              movieID,
		Title:           title,
		Description:     req.Description,
		DurationMinutes: req.DurationMinutes,
		Genre:           req.Genre,
		AgeRating:       req.AgeRating,
	}

	err = s.movieRepo.Update(
		ctx,
		movie,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}

		return nil, err
	}

	return movie, nil
}

func (s *MovieService) Delete(
	ctx context.Context,
	id string,
) error {
	movieID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("id movie tidak valid")
	}

	return s.movieRepo.Delete(
		ctx,
		movieID,
	)
}
