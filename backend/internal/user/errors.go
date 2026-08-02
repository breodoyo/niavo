package user

import "errors"

var (
	ErrUserNotFound      = errors.New("User not found")
	ErrEmailAreadyExists = errors.New("Email already in use")
)
