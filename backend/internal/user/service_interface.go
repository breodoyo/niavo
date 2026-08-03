package user

import "context"

type Service interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, id string) (User, error)
	UpdateUser(ctx context.Context, id string, req UpdateUserRequest) (User, error)
	DeleteUser(ctx context.Context, id string) error
}