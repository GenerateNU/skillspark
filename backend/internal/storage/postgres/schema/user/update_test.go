package user

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Update(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewUserRepository(testDB)
	ctx := context.Background()

	// Setup: Create a user
	userInput := &models.CreateUserInput{}
	userInput.Body.Name = "Original Name"
	userInput.Body.Email = "original@test.com"
	userInput.Body.Username = "original"
	authID := uuid.New()
	userInput.Body.AuthID = authID

	createdUser, err := repo.CreateUser(ctx, userInput)
	if err != nil {
		t.Fatalf("Failed to create user setup: %v", err)
	}

	t.Run("Success Partial Update", func(t *testing.T) {
		newName := "Updated Name"
		updateInput := &models.UpdateUserInput{}
		updateInput.ID = createdUser.ID
		updateInput.Body.Name = &newName

		updatedUser, err := repo.UpdateUser(ctx, updateInput)
		assert.NoError(t, err)
		assert.NotNil(t, updatedUser)
		assert.Equal(t, newName, updatedUser.Name)
		assert.Equal(t, createdUser.Email, updatedUser.Email) // Should remain unchanged
	})

	t.Run("Not Found", func(t *testing.T) {
		randomID := uuid.New()
		newName := "Ghost User"
		updateInput := &models.UpdateUserInput{}
		updateInput.ID = randomID
		updateInput.Body.Name = &newName

		updatedUser, err := repo.UpdateUser(ctx, updateInput)
		assert.Error(t, err)
		assert.Nil(t, updatedUser)
	})
}
