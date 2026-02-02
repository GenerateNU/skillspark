package repomocks

import (
	"context"
	"skillspark/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *models.CreateUserInput) (*models.User, error) {
	args := m.Called(ctx, user)
	var u *models.User
	if v := args.Get(0); v != nil {
		u = v.(*models.User)
	}
	return u, args.Error(1)
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	args := m.Called(ctx, id)
	var u *models.User
	if v := args.Get(0); v != nil {
		u = v.(*models.User)
	}
	return u, args.Error(1)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user *models.UpdateUserInput) (*models.User, error) {
	args := m.Called(ctx, user)
	var u *models.User
	if v := args.Get(0); v != nil {
		u = v.(*models.User)
	}
	return u, args.Error(1)
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	args := m.Called(ctx, id)
	var u *models.User
	if v := args.Get(0); v != nil {
		u = v.(*models.User)
	}
	return u, args.Error(1)
}
