package workflow

import "testing"

func TestCreateWorkflowRequestValidate(t *testing.T) {
	tests := []struct {
		name    string
		request CreateWorkflowRequest
		wantErr bool
	}{
		{
			name:    "valid workflow with description",
			request: CreateWorkflowRequest{Name: "Support Tickets", Description: "Track requests"},
			wantErr: false,
		},
		{
			name:    "valid workflow without description",
			request: CreateWorkflowRequest{Name: "Onboarding"},
			wantErr: false,
		},
		{
			name:    "missing name",
			request: CreateWorkflowRequest{Name: "", Description: "Something"},
			wantErr: true,
		},
		{
			name:    "spaces only name",
			request: CreateWorkflowRequest{Name: "   "},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}