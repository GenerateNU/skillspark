package user

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_GetByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewUserRepository(testDB)
	ctx := context.Background()

	// Setup: Create a user
	userInput := &models.CreateUserInput{}
	userInput.Body.Name = "GetByID User"
	userInput.Body.Email = "getbyid@test.com"
	userInput.Body.Username = "getbyid"
	authID := uuid.New()
	userInput.Body.AuthID = &authID

	createdUser, err := repo.CreateUser(ctx, userInput)
	if err != nil {
		t.Fatalf("Failed to create user setup: %v", err)
	}

	t.Run("Success", func(t *testing.T) {
		user, err := repo.GetUserByID(ctx, createdUser.ID)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, createdUser.ID, user.ID)
		assert.Equal(t, createdUser.Email, user.Email)
	})

	t.Run("Not Found", func(t *testing.T) {
		randomID := uuid.New()
		user, err := repo.GetUserByID(ctx, randomID)
		assert.Error(t, err)
		assert.Nil(t, user)
		// Assuming the error is typed or contains specific message, but generic error check is fine for now
		// If we want to be specific, we'd check for custom error type, but standard assert.Error is safe.
	})
}
