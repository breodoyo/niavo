package user

import "context"

type MockUserService struct {}

func (m *MockUserService) CreateUser(ctx context.Context, req CreateUserRequest) (User, error) {
	return User{}, nil
}
func (m *MockUserService) ListUsers(ctx context.Context) ([]User, error) {
	return []User{}, nil
}
func (m *MockUserService) GetUser(ctx context.Context, id string) (User, error) {
	return User{
		ID: id,
		Email: "ksks@gmail.com",
		FirstName: "Nan",
		LastName: "Luis",
	}, nil
}