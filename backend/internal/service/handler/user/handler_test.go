package user

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_GetUserByUsername(t *testing.T) {
	tests := []struct {
		name      string
		username  string
		mockSetup func(*repomocks.MockUserRepository)
		wantErr   bool
	}{
		{
			name:     "successful get user by username",
			username: "jamesw",
			mockSetup: func(m *repomocks.MockUserRepository) {
				m.On("GetUserByUsername", mock.Anything, "jamesw").
					Return(&models.User{
						ID:                 uuid.MustParse("b8c9d0e1-f2a3-4b4c-5d6e-7f8a9b0c1d2e"),
						Name:               "James Wilson",
						Email:              "james.wilson@email.com",
						Username:           "jamesw",
						LanguagePreference: "en",
						AuthID:             uuid.MustParse("00000000-0000-0000-0000-000000000002"),
						CreatedAt:          time.Now(),
						UpdatedAt:          time.Now(),
					}, nil)
			},
			wantErr: false,
		},
		{
			name:     "username not found",
			username: "randomusername",
			mockSetup: func(m *repomocks.MockUserRepository) {
				m.On("GetUserByUsername", mock.Anything, "randomusername").
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("User", "username", "randomusername").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockRepo := new(repomocks.MockUserRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()
			input := &models.GetUserByUsernameInput{Username: tt.username}

			user, err := handler.GetUserByUsername(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.username, user.Username)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
