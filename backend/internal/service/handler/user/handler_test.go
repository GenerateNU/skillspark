package user

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_GetUserByUsername(t *testing.T) {
	tests := []struct {
		name      string
		username  string
		mockSetup func(*repomocks.MockUserRepository)
		wantErr   bool
		wantExist bool
	}{
		{
			name:     "username exists",
			username: "jamesw",
			mockSetup: func(m *repomocks.MockUserRepository) {
				m.On("GetUserByUsername", mock.Anything, "jamesw").Return(true, nil)
			},
			wantErr:   false,
			wantExist: true,
		},
		{
			name:     "username does not exist",
			username: "randomusername",
			mockSetup: func(m *repomocks.MockUserRepository) {
				m.On("GetUserByUsername", mock.Anything, "randomusername").Return(false, nil)
			},
			wantErr:   false,
			wantExist: false,
		},
		{
			name:     "repository error",
			username: "erroruser",
			mockSetup: func(m *repomocks.MockUserRepository) {
				m.On("GetUserByUsername", mock.Anything, "erroruser").Return(false, &errs.HTTPError{
					Code:    500,
					Message: "internal server error",
				})
			},
			wantErr:   true,
			wantExist: false,
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

			exists, err := handler.GetUserByUsername(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.False(t, exists)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantExist, exists)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
