package user

import (
	"context"

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
		Email: req.Email,
		FirstName: req.FirstName,
		LastName: req.LastName,
		PasswordHash: string(hashedPassword),
	}
	return s.repo.CreateUser(ctx, user)
}
func (s *UserService) ListUsers(ctx context.Context) ([]User, error) {
	return s.repo.ListUsers(ctx)
}
func (s *UserService) GetUser(ctx context.Context, id string) (User, error) {
	return s.repo.GetUser(ctx, id)
}