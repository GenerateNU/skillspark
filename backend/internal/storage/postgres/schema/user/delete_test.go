package user

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Delete(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewUserRepository(testDB)
	ctx := context.Background()

	// Setup: Create a user
	userInput := &models.CreateUserInput{}
	userInput.Body.Name = "Delete User"
	userInput.Body.Email = "delete@test.com"
	userInput.Body.Username = "deleteuser"
	authID := uuid.New()
	userInput.Body.AuthID = authID

	createdUser, err := repo.CreateUser(ctx, userInput)
	if err != nil {
		t.Fatalf("Failed to create user setup: %v", err)
	}

	t.Run("Success", func(t *testing.T) {
		deletedUser, err := repo.DeleteUser(ctx, createdUser.ID)
		assert.NoError(t, err)
		assert.NotNil(t, deletedUser)
		assert.Equal(t, createdUser.ID, deletedUser.ID)

		// Verify it's gone
		_, err = repo.GetUserByID(ctx, createdUser.ID)
		assert.Error(t, err)
	})

	t.Run("Not Found", func(t *testing.T) {
		randomID := uuid.New()
		deletedUser, err := repo.DeleteUser(ctx, randomID)
		assert.Error(t, err)
		assert.Nil(t, deletedUser)
	})
}
