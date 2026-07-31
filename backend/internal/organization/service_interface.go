package organization

import (
	"context"
)

type OrganizationService interface {
	CreateOrganization(
		ctx context.Context,
		org Organization,
	) (Organization, error)

	ListOrganizations(
		ctx context.Context,
	) ([]Organization, error)

	GetOrganization(
		ctx context.Context,
		id string,
	) (Organization, error)

	UpdateOrganization(
		ctx context.Context,
		id string,
		name string,
	) (Organization, error)

	DeleteOrganization(
		ctx context.Context,
		id string,
	) error
}