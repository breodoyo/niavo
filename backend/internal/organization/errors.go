package organization

import "errors"

var (
	ErrInvalidName          = errors.New("organization name is required")
	ErrInvalidSlug          = errors.New("organization slug is required")
	ErrOrganizationNotFound = errors.New("organization not found")
)
