package organization

import "context"

type MockOrganizationService struct {}

func (m MockOrganizationService) CreateOrganization(
	ctx context.Context,
	org Organization,
) (Organization, error) {
	org.ID = "123"
	return org, nil
}
func (m *MockOrganizationService) ListOrganizations(ctx context.Context) ([]Organization, error) {
	return []Organization{}, nil
}

func (m *MockOrganizationService) GetOrganization(ctx context.Context, id string) (Organization, error) {
	return Organization{
		ID:   id,
		Name: "Niavo",
		Slug: "niavo",
	}, nil
}

func (m *MockOrganizationService) UpdateOrganization(ctx context.Context, id, name string) (Organization, error) {
	return Organization{
		ID:   id,
		Name: name,
		Slug: "niavo",
	}, nil
}

func (m *MockOrganizationService) DeleteOrganization(ctx context.Context, id string) error {
	return nil
}