package auth

import (
	"context"
	"errors"

	"github.com/breodoyo/niavo/backend/internal/user"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

type Service struct {
	userService user.Service
	jwtSecret   string
}

func NewService(userService user.Service, jwtSecret string) *Service {
	return &Service{
		userService: userService,
		jwtSecret:   jwtSecret,
	}
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (string, error) {
	u, err := s.userService.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(u.PasswordHash),
		[]byte(req.Password),
	); err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := GenerateToken(u.ID, s.jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}