package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	authLib "skillspark/internal/auth"
	"skillspark/internal/config"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRoundTripper is a helper to mock http.Client via its Transport
type MockRoundTripper struct {
	RoundTripFunc func(req *http.Request) *http.Response
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req), nil
}

func setupMockAuthClient(t *testing.T, responseBody interface{}, statusCode int) {
	originalClient := authLib.Client
	t.Cleanup(func() {
		authLib.Client = originalClient
	})

	mockTransport := &MockRoundTripper{
		RoundTripFunc: func(req *http.Request) *http.Response {
			respBytes, _ := json.Marshal(responseBody)
			return &http.Response{
				StatusCode: statusCode,
				Body:       io.NopCloser(bytes.NewBuffer(respBytes)),
				Header:     make(http.Header),
			}
		},
	}

	authLib.Client = &http.Client{
		Transport: mockTransport,
		Timeout:   1 * time.Second,
	}
}

func TestHandler_GuardianLogin(t *testing.T) {
	tests := []struct {
		name          string
		input         *models.LoginInput
		mockSetup     func(*repomocks.MockGuardianRepository)
		authResponse  interface{}
		authStatus    int
		wantErr       bool
		expectedError string
	}{
		{
			name: "successful login",
			input: &models.LoginInput{
				Body: struct {
					Email    string `json:"email" db:"email"`
					Password string `json:"password" db:"password"`
				}{
					Email:    "guardian@example.com",
					Password: "password123",
				},
			},
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("GetGuardianByAuthID", mock.Anything, "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11").Return(&models.Guardian{
					ID:     uuid.New(),
					UserID: uuid.New(),
				}, nil)
			},
			authResponse: models.LoginResponse{
				AccessToken: "test-token",
				ExpiresIn:   3600,
				User: models.UserResponse{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000000"), // We'll use a string for ID in logic but model has UUID
				},
			},
			authStatus: http.StatusOK,
			wantErr:    false,
		},
		{
			name: "supabase login failure",
			input: &models.LoginInput{
				Body: struct {
					Email    string `json:"email" db:"email"`
					Password string `json:"password" db:"password"`
				}{
					Email:    "fail@example.com",
					Password: "wrongpassword",
				},
			},
			mockSetup: func(m *repomocks.MockGuardianRepository) {
			},
			authResponse: map[string]interface{}{
				"error":             "invalid_grant",
				"error_description": "Invalid login credentials",
			},
			authStatus:    http.StatusBadRequest,
			wantErr:       true,
			expectedError: "400",
		},
		{
			name: "guardian not found",
			input: &models.LoginInput{
				Body: struct {
					Email    string `json:"email" db:"email"`
					Password string `json:"password" db:"password"`
				}{
					Email:    "new@example.com",
					Password: "password123",
				},
			},
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				notFoundErr := errs.NotFound("Guardian", "auth_id", "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22")
				m.On("GetGuardianByAuthID", mock.Anything, "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22").Return(nil, &notFoundErr)
			},
			// Note: modifying authResponse logic below to return correct ID for this case
			authResponse: models.LoginResponse{
				AccessToken: "test-token",
				User: models.UserResponse{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"), // different ID to match logic if needed
				},
			},
			authStatus:    http.StatusOK,
			wantErr:       true,
			expectedError: "not found",
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel
		t.Run(tt.name, func(t *testing.T) {

			// Fixup for specific tests
			if tt.name == "successful login" {
				resp := tt.authResponse.(models.LoginResponse)
				resp.User.ID = uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11") // Use valid UUID
				tt.authResponse = resp
			}
			if tt.name == "guardian not found" {
				resp := tt.authResponse.(models.LoginResponse)
				resp.User.ID = uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22")
				tt.authResponse = resp
			}

			setupMockAuthClient(t, tt.authResponse, tt.authStatus)

			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockUserRepo := new(repomocks.MockUserRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			tt.mockSetup(mockGuardianRepo)

			// Config mock
			cfg := config.Supabase{
				URL:            "http://mock-supabase",
				ServiceRoleKey: "mock-key",
			}

			handler := NewHandler(cfg, mockUserRepo, mockGuardianRepo, mockManagerRepo)
			ctx := context.Background()

			output, err := handler.GuardianLogin(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedError != "" {
					assert.Contains(t, err.Error(), tt.expectedError)
				}
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, "test-token", output.AccessTokenCookie.Value)
			}
			mockGuardianRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_ManagerLogin(t *testing.T) {
	tests := []struct {
		name          string
		input         *models.LoginInput
		mockSetup     func(*repomocks.MockManagerRepository)
		authResponse  interface{}
		authStatus    int
		wantErr       bool
		expectedError string
	}{
		{
			name: "successful manager login",
			input: &models.LoginInput{
				Body: struct {
					Email    string `json:"email" db:"email"`
					Password string `json:"password" db:"password"`
				}{
					Email:    "manager@example.com",
					Password: "password123",
				},
			},
			mockSetup: func(m *repomocks.MockManagerRepository) {
				// We'll fixup ID inside the loop same as GuardianLogin if expected
				m.On("GetManagerByAuthID", mock.Anything, "manager-auth-id").Return(&models.Manager{
					ID: uuid.New(),
				}, nil)
			},
			authResponse: models.LoginResponse{
				AccessToken: "manager-token",
				ExpiresIn:   3600,
				User: models.UserResponse{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000000"), // placeholder
				},
			},
			authStatus: http.StatusOK,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "successful manager login" {
				resp := tt.authResponse.(models.LoginResponse)
				resp.User.ID = uuid.MustParse("b0eebc99-9c0b-4ef8-bb6d-6bb9bd380b22")
				tt.authResponse = resp
				tt.mockSetup = func(m *repomocks.MockManagerRepository) {
					m.On("GetManagerByAuthID", mock.Anything, "b0eebc99-9c0b-4ef8-bb6d-6bb9bd380b22").Return(&models.Manager{
						ID: uuid.New(),
					}, nil)
				}
			}

			setupMockAuthClient(t, tt.authResponse, tt.authStatus)

			mockManagerRepo := new(repomocks.MockManagerRepository)
			// Need dummy repos for NewHandler
			mockUserRepo := new(repomocks.MockUserRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)

			tt.mockSetup(mockManagerRepo)

			cfg := config.Supabase{URL: "http://mock", ServiceRoleKey: "key"}
			handler := NewHandler(cfg, mockUserRepo, mockGuardianRepo, mockManagerRepo)

			output, err := handler.ManagerLogin(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, "manager-token", output.AccessTokenCookie.Value)
			}
			mockManagerRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GuardianSignUp(t *testing.T) {
	tests := []struct {
		name          string
		input         *models.GuardianSignUpInput
		mockSetup     func(*repomocks.MockGuardianRepository, *repomocks.MockUserRepository)
		authResponse  interface{}
		authStatus    int
		wantErr       bool
		expectedError string
	}{
		{
			name: "successful signup",
			input: &models.GuardianSignUpInput{
				Body: struct {
					Name                string  `json:"name" db:"name"`
					Email               string  `json:"email" db:"email"`
					Username            string  `json:"username" db:"username"`
					Password            string  `json:"password" db:"password"`
					ProfilePictureS3Key *string `json:"profile_picture_s3_key" db:"profile_picture_s3_key"`
					LanguagePreference  string  `json:"language_preference" db:"language_preference"`
				}{
					Name:               "Guardian Name",
					Email:              "signup@example.com",
					Username:           "guser",
					Password:           "StrongPass1!",
					LanguagePreference: "en",
				},
			},
			mockSetup: func(mg *repomocks.MockGuardianRepository, mu *repomocks.MockUserRepository) {
				mg.On("CreateGuardian", mock.Anything, mock.AnythingOfType("*models.CreateGuardianInput")).Return(&models.Guardian{
					ID:     uuid.New(),
					UserID: uuid.New(),
				}, nil)
				mu.On("UpdateUser", mock.Anything, mock.AnythingOfType("*models.UpdateUserInput")).Return(nil, nil)
			},
			authResponse: models.SignupResponse{
				AccessToken: "signup-token",
				User: models.UserSignupResponse{
					ID: uuid.MustParse("c0eebc99-9c0b-4ef8-bb6d-6bb9bd380c33"),
				},
			},
			authStatus: http.StatusOK,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMockAuthClient(t, tt.authResponse, tt.authStatus)

			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockUserRepo := new(repomocks.MockUserRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)

			tt.mockSetup(mockGuardianRepo, mockUserRepo)

			cfg := config.Supabase{URL: "http://mock", ServiceRoleKey: "key"}
			handler := NewHandler(cfg, mockUserRepo, mockGuardianRepo, mockManagerRepo)

			output, err := handler.GuardianSignUp(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, "signup-token", output.Body.Token)
			}
			mockGuardianRepo.AssertExpectations(t)
			mockUserRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_ManagerSignUp(t *testing.T) {
	tests := []struct {
		name          string
		input         *models.ManagerSignUpInput
		mockSetup     func(*repomocks.MockManagerRepository, *repomocks.MockUserRepository)
		authResponse  interface{}
		authStatus    int
		wantErr       bool
		expectedError string
	}{
		{
			name: "successful manager signup",
			input: &models.ManagerSignUpInput{
				Body: struct {
					Name                string     `json:"name" db:"name" doc:"name of the manager" required:"true"`
					Email               string     `json:"email" db:"email" doc:"email of the manager" required:"true"`
					Username            string     `json:"username" db:"username" doc:"username of the manager" required:"true"`
					Password            string     `json:"password" db:"password" doc:"password of the manager" required:"true"`
					ProfilePictureS3Key *string    `json:"profile_picture_s3_key" db:"profile_picture_s3_key" doc:"profile picture s3 key of the manager" required:"false"`
					LanguagePreference  string     `json:"language_preference" db:"language_preference" doc:"language preference of the manager" required:"false"`
					OrganizationID      *uuid.UUID `json:"organization_id" db:"organization_id" doc:"organization id of the organization the manager is associated with" required:"false"`
					Role                string     `json:"role" db:"role" doc:"role of the manager being created" required:"false"`
				}{
					Name:               "Manager Name",
					Email:              "msignup@example.com",
					Username:           "muser",
					Password:           "StrongPass1!",
					LanguagePreference: "en",
				},
			},
			mockSetup: func(mm *repomocks.MockManagerRepository, mu *repomocks.MockUserRepository) {
				mm.On("CreateManager", mock.Anything, mock.AnythingOfType("*models.CreateManagerInput")).Return(&models.Manager{
					ID:     uuid.New(),
					UserID: uuid.New(),
				}, nil)
				mu.On("UpdateUser", mock.Anything, mock.AnythingOfType("*models.UpdateUserInput")).Return(nil, nil)
			},
			authResponse: models.SignupResponse{
				AccessToken: "msignup-token",
				User: models.UserSignupResponse{
					ID: uuid.MustParse("d0eebc99-9c0b-4ef8-bb6d-6bb9bd380d44"),
				},
			},
			authStatus: http.StatusOK,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMockAuthClient(t, tt.authResponse, tt.authStatus)

			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockUserRepo := new(repomocks.MockUserRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)

			tt.mockSetup(mockManagerRepo, mockUserRepo)

			cfg := config.Supabase{URL: "http://mock", ServiceRoleKey: "key"}
			handler := NewHandler(cfg, mockUserRepo, mockGuardianRepo, mockManagerRepo)

			output, err := handler.ManagerSignUp(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, "msignup-token", output.Body.Token)
			}
			mockManagerRepo.AssertExpectations(t)
			mockUserRepo.AssertExpectations(t)
		})
	}
}
