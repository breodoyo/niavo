package user

import "testing"

func TestCreateUserRequestValidate(t *testing.T) {
	tests := []struct {
		name    string
		request CreateUserRequest
		wantErr bool
	}{
		{
			name: "valid user",
			request: CreateUserRequest{
				Email: "bre04@gmail.com",
				FirstName: "Bree",
				LastName: "Johns",
				Password: "Hello123!",
			},
			wantErr: false,
		},
		{
			name: "missing name",
			request: CreateUserRequest{
				FirstName: "",
				LastName: "Johns",
			},
			wantErr: true,
		},
		{
			name: "missing email",
			request: CreateUserRequest{
				Email: "",
				FirstName: "Bree",
				LastName:  "Johns",
			},
			wantErr: true,
		},
		{
			name: "malformed_email",
			request: CreateUserRequest{
				Email: "bre@gmailcom",
				FirstName: "Bree",
				LastName: "Johns",
			},
			wantErr: true,
		},
		{
			name: "spaces only first_name",
			request: CreateUserRequest{
				FirstName: "   ",
				LastName: "Johns",
			},
			wantErr: true,
		},
		{
			name: "valid atleast 8 characters",
			request: CreateUserRequest{
				Email: "bre@gmail.com",
				FirstName: "Bree",
				LastName: "Johns",
				Password: "Hello123!",
			},
			wantErr: false,
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