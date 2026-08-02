package user

import (
	"errors"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

type CreateUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func (r CreateUserRequest) Validate() error {
	if strings.TrimSpace(r.FirstName) == "" {
		return errors.New("firstname is required")
	}
	if strings.TrimSpace(r.LastName) == "" {
		return errors.New("lastname is required")
	}
	if strings.TrimSpace(r.Email) == "" {
		return errors.New("email is require")
	}
	if !emailRegex.MatchString(r.Email) {
		return errors.New("email format is invalid")
	}
	if len(r.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}
