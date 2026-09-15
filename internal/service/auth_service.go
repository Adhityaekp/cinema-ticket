package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/Adhityaekp/cinema-ticket/config"
	"github.com/Adhityaekp/cinema-ticket/internal/dto"
	"github.com/Adhityaekp/cinema-ticket/internal/model"
	"github.com/Adhityaekp/cinema-ticket/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo     *repository.UserRepository
	emailService *EmailService
	cfg          *config.Config
}

func NewAuthService(
	userRepo *repository.UserRepository,
	emailService *EmailService,
	cfg *config.Config,
) *AuthService {

	return &AuthService{
		userRepo:     userRepo,
		emailService: emailService,
		cfg:          cfg,
	}
}

func (s *AuthService) Register(
	ctx context.Context,
	req dto.RegisterRequest,
) (*model.User, error) {

	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	existingUser, err := s.userRepo.FindByEmail(ctx, req.Email)

	if err == nil && existingUser != nil {
		return nil, errors.New("email sudah terdaftar")
	}

	if err != nil && !errors.Is(err, context.Canceled) {
		// lanjut jika user memang tidak ditemukan
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	verificationToken, err := generateVerificationToken()
	if err != nil {
		return nil, err
	}

	expiredAt := time.Now().Add(30 * time.Minute)

	user := &model.User{
		Name:                       req.Name,
		Email:                      req.Email,
		Password:                   string(hashedPassword),
		Role:                       "CUSTOMER",
		IsActive:                   true,
		EmailVerificationToken:     &verificationToken,
		EmailVerificationExpiresAt: &expiredAt,
	}

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	err = s.emailService.SendVerificationEmail(
		user.Email,
		verificationToken,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) VerifyEmail(
	ctx context.Context,
	token string,
) error {

	user, err := s.userRepo.FindByVerificationToken(ctx, token)
	if err != nil {
		return errors.New("token verifikasi tidak valid")
	}

	if user.EmailVerifiedAt != nil {
		return errors.New("email sudah diverifikasi")
	}

	if user.EmailVerificationExpiresAt == nil ||
		time.Now().After(*user.EmailVerificationExpiresAt) {

		return errors.New("token verifikasi sudah expired")
	}

	return s.userRepo.VerifyEmail(ctx, user.ID)
}

func (s *AuthService) Login(
	ctx context.Context,
	req dto.LoginRequest,
) (*model.User, error) {

	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	if !user.IsActive {
		return nil, errors.New("akun tidak aktif")
	}

	if user.EmailVerifiedAt == nil {
		return nil, errors.New("email belum diverifikasi")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	return user, nil
}

func generateVerificationToken() (string, error) {

	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
