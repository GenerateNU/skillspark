package user

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_GetByUsername(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewUserRepository(testDB)
	ctx := context.Background()

	// Setup: Create a user
	userInput := &models.CreateUserInput{}
	userInput.Body.Name = "GetByUsername User"
	userInput.Body.Email = "getbyusername@test.com"
	userInput.Body.Username = "getbyusername"
	authID := uuid.New()
	userInput.Body.AuthID = authID

	createdUser, err := repo.CreateUser(ctx, userInput)
	if err != nil {
		t.Fatalf("Failed to create user setup: %v", err)
	}

	t.Run("Success", func(t *testing.T) {
		user, err := repo.GetUserByUsername(ctx, createdUser.Username)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, createdUser.ID, user.ID)
		assert.Equal(t, createdUser.Username, user.Username)
		assert.Equal(t, createdUser.Email, user.Email)
	})

	t.Run("Not Found", func(t *testing.T) {
		user, err := repo.GetUserByUsername(ctx, "randomusername")
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}
