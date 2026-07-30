package organization

import (
	"errors"
	"strings"
)

type CreateOrganizationRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type UpdateOrganizationRequest struct {
	Name string `json:"name"`
}

func (r CreateOrganizationRequest) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("organization name is required")
	}

	if strings.TrimSpace(r.Slug) == "" {
		return errors.New("organization slug is required")
	}

	return nil
}

func (r UpdateOrganizationRequest) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("organization name is required")
	}

	return nil
}