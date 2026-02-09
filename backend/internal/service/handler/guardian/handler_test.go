package guardian

import (
	"context"
	"skillspark/internal/config"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	repomocks "skillspark/internal/storage/repo-mocks"
	"testing"
	"time"
	"net/http"
	supabaseMock "skillspark/internal/service/handler/auth"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_GetGuardianById(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(*repomocks.MockGuardianRepository)
		wantErr   bool
	}{
		{
			name: "successful get guardian by id - Sarah Johnson",
			id:   "11111111-1111-1111-1111-111111111111",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("GetGuardianByID", mock.Anything, uuid.MustParse("11111111-1111-1111-1111-111111111111")).Return(&models.Guardian{
					ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					UserID:    uuid.MustParse("a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d"),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "guardian not found",
			id:   "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("GetGuardianByID", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Guardian", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
		{
			name: "internal server error",
			id:   "ffffffff-ffff-ffff-ffff-ffffffffffff",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("GetGuardianByID", mock.Anything, uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff")).
					Return(nil, &errs.HTTPError{
						Code:    errs.InternalServerError("Internal server error").Code,
						Message: "Internal server error",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			cfg := config.Supabase{
				URL:            "http://mock-supabase",
				ServiceRoleKey: "mock-key",
			}

			testDB := testutil.SetupTestDB(t)

			handler := NewHandler(mockRepo, testDB, cfg)
			ctx := context.Background()

			input := &models.GetGuardianByIDInput{ID: uuid.MustParse(tt.id)}
			guardian, err := handler.GetGuardianById(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, guardian)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, guardian)
				assert.Equal(t, tt.id, guardian.ID.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateGuardian(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.CreateGuardianInput
		mockSetup func(*repomocks.MockGuardianRepository, )
		wantErr   bool
	}{
		{
			name: "successful create guardian - Michael Chen",
			input: func() *models.CreateGuardianInput {
				input := &models.CreateGuardianInput{}
				input.Body.Name = "Michael Chen"
				input.Body.Email = "michael.chen@example.com"
				input.Body.Username = "mchen"
				input.Body.LanguagePreference = "en"
				return input
			}(),
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("CreateGuardian", mock.Anything, mock.AnythingOfType("*models.CreateGuardianInput")).Return(&models.Guardian{
					ID:        uuid.New(),
					UserID:    uuid.New(),
					Name:      "Michael Chen",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "internal server error during creation",
			input: func() *models.CreateGuardianInput {
				input := &models.CreateGuardianInput{}
				input.Body.Name = "Error User"
				input.Body.Email = "error@example.com"
				input.Body.Username = "error"
				input.Body.LanguagePreference = "en"
				return input
			}(),
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("CreateGuardian", mock.Anything, mock.AnythingOfType("*models.CreateGuardianInput")).
					Return(nil, &errs.HTTPError{
						Code:    errs.InternalServerError("Internal server error").Code,
						Message: "Internal server error",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			cfg := config.Supabase{
				URL:            "http://mock-supabase",
				ServiceRoleKey: "mock-key",
			}

			testDB := testutil.SetupTestDB(t)

			handler := NewHandler(mockRepo, testDB, cfg)
			ctx := context.Background()

			guardian, err := handler.CreateGuardian(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, guardian)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, guardian)
				assert.Equal(t, tt.input.Body.Name, guardian.Name)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_UpdateGuardian(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		input     *models.UpdateGuardianInput
		mockSetup func(*repomocks.MockGuardianRepository)
		wantErr   bool
	}{
		{
			name: "successful update guardian - Sarah Johnson",
			id:   "11111111-1111-1111-1111-111111111111",
			input: func() *models.UpdateGuardianInput {
				input := &models.UpdateGuardianInput{}
				input.ID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
				input.Body.Name = "Sarah Johnson"
				input.Body.Email = "sarah.j@example.com"
				input.Body.Username = "sjohnson"
				input.Body.LanguagePreference = "es"
				return input
			}(),
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("UpdateGuardian", mock.Anything, mock.MatchedBy(func(input *models.UpdateGuardianInput) bool {
					return input.ID == uuid.MustParse("11111111-1111-1111-1111-111111111111")
				})).Return(&models.Guardian{
					ID:                 uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					UserID:             uuid.New(),
					Name:               "Sarah Johnson",
					Email:              "sarah.j@example.com",
					Username:           "sjohnson",
					LanguagePreference: "es",
					CreatedAt:          time.Now(),
					UpdatedAt:          time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "guardian not found",
			id:   "00000000-0000-0000-0000-000000000000",
			input: func() *models.UpdateGuardianInput {
				input := &models.UpdateGuardianInput{}
				input.ID = uuid.MustParse("00000000-0000-0000-0000-000000000000")
				input.Body.Name = "Unknown"
				return input
			}(),
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("UpdateGuardian", mock.Anything, mock.AnythingOfType("*models.UpdateGuardianInput")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Guardian", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			cfg := config.Supabase{
				URL:            "http://mock-supabase",
				ServiceRoleKey: "mock-key",
			}

			testDB := testutil.SetupTestDB(t)

			handler := NewHandler(mockRepo, testDB, cfg)
			ctx := context.Background()

			guardian, err := handler.UpdateGuardian(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, guardian)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, guardian)
				assert.Equal(t, tt.input.Body.Name, guardian.Name)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetGuardianByChildId(t *testing.T) {
	tests := []struct {
		name      string
		childID   string
		mockSetup func(*repomocks.MockGuardianRepository)
		wantErr   bool
	}{
		{
			name:    "successful get guardian by child id - Emily Johnson",
			childID: "30000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("GetGuardianByChildID", mock.Anything, uuid.MustParse("30000000-0000-0000-0000-000000000001")).Return(&models.Guardian{
					ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					UserID:    uuid.MustParse("a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d"),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name:    "guardian not found for child",
			childID: "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("GetGuardianByChildID", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Guardian", "child_id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			cfg := config.Supabase{
				URL:            "http://mock-supabase",
				ServiceRoleKey: "mock-key",
			}

			testDB := testutil.SetupTestDB(t)

			handler := NewHandler(mockRepo, testDB, cfg)
			ctx := context.Background()

			input := &models.GetGuardianByChildIDInput{ChildID: uuid.MustParse(tt.childID)}
			guardian, err := handler.GetGuardianByChildId(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, guardian)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, guardian)
				assert.Equal(t, "11111111-1111-1111-1111-111111111111", guardian.ID.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_DeleteGuardian(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(*repomocks.MockGuardianRepository)
		authResponse  interface{}
		authStatus    int
		wantErr   bool
	}{
		{
			name: "successful delete guardian", 
			id:   "761ef221-6a5a-463e-8b1f-a3a9296c7fb9",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("DeleteGuardian", mock.Anything, uuid.MustParse("761ef221-6a5a-463e-8b1f-a3a9296c7fb9")).Return(&models.Guardian{
					ID:        uuid.MustParse("761ef221-6a5a-463e-8b1f-a3a9296c7fb9"),
					UserID:    uuid.MustParse("484de30a-aaa3-4a3a-aeb7-14d7f7ddbe26"),
				}, nil)
			},
			authResponse: []string {},
			authStatus: http.StatusOK,
			wantErr: false,
		},
		{
			name: "guardian not found",
			id:   "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("DeleteGuardian", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Guardian", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			authResponse: []string {},
			authStatus: http.StatusOK,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			supabaseMock.SetupMockAuthClient(t, tt.authResponse, tt.authStatus)

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			cfg := config.Supabase{
				URL:            "http://mock",
				ServiceRoleKey: "key",
			}

			testDB := testutil.SetupTestDB(t)

			handler := NewHandler(mockRepo, testDB, cfg)
			ctx := context.Background()

			input := &models.DeleteGuardianInput{ID: uuid.MustParse(tt.id)}
			guardian, err := handler.DeleteGuardian(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, guardian)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, guardian)
				assert.Equal(t, tt.id, guardian.ID.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
