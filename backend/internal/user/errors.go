package user

import "errors"

var (
	ErrUserNotFound      = errors.New("User not found")
	ErrEmailAlreadyExists = errors.New("Email already in use")
)
