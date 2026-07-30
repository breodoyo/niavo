package organization

import "testing"

func TestCreateOrganizationRequestValidate(t *testing.T) {
	tests := []struct {
		name    string
		request CreateOrganizationRequest
		wantErr bool
	}{
		{
			name: "valid organization",
			request: CreateOrganizationRequest{
				Name: "Niavo",
				Slug: "niavo",
			},
			wantErr: false,
		},
		{
			name: "missing name",
			request: CreateOrganizationRequest{
				Name: "",
				Slug: "niavo",
			},
			wantErr: true,
		},
		{
			name: "missing slug",
			request: CreateOrganizationRequest{
				Name: "Niavo",
				Slug: "",
			},
			wantErr: true,
		},
		{
			name: "spaces only name",
			request: CreateOrganizationRequest{
				Name: "   ",
				Slug: "niavo",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf(
					"Validate() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}