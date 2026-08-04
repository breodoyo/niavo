package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *Repository
}

func NewService(repo *Repository) *UserService {
	return &UserService{repo: repo}
}
func (s *UserService) CreateUser(ctx context.Context, req CreateUserRequest) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return User{}, err
	}
	user := User{
		Email:        req.Email,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PasswordHash: string(hashedPassword),
	}
	created, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return User{}, ErrEmailAlreadyExists
		}
		return User{}, err
	}
	return created, nil
}
func (s *UserService) ListUsers(ctx context.Context) ([]User, error) {
	return s.repo.ListUsers(ctx)
}
func (s *UserService) GetUser(ctx context.Context, id string) (User, error) {
	return s.repo.GetUser(ctx, id)
}
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}
func (s *UserService) UpdateUser(ctx context.Context, id string, req UpdateUserRequest) (User, error) {
	return s.repo.UpdateUser(ctx, id, req.FirstName, req.LastName)
}
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}
