package organization

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) CreateOrganization(
	ctx context.Context,
	org Organization,
) (Organization, error) {

	if org.Name == "" {
		return Organization{}, ErrInvalidName
	}
	if org.Slug == "" {
		return Organization{}, ErrInvalidSlug
	}
	return s.repo.CreateOrganization(ctx, org)
}
func (s *Service) ListOrganizations(
	ctx context.Context,
) ([]Organization, error) {
	return s.repo.ListOrganizations(ctx)
}
func (s *Service) GetOrganization(
	ctx context.Context,
	id string,
) (Organization, error) {
	return s.repo.GetOrganization(ctx, id)
}
func (s *Service) UpdateOrganization(
	ctx context.Context,
	id string,
	name string,
) (Organization, error) {
	if name == "" {
		return Organization{}, ErrInvalidName
	}
	return s.repo.UpdateOrganization(ctx, id, name)
}
func (s *Service) DeleteOrganization(
	ctx context.Context,
	id string,
) error {
	return s.repo.DeleteOrganization(ctx, id)
}