package guardian

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGuardianRepository_Create_David_Kim(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()

	guardianInput := func() *models.CreateGuardianInput {
		input := &models.CreateGuardianInput{}
		input.Body.Name = "David Kim"
		input.Body.Email = "david.kim@test.com"
		input.Body.Username = "davidk"
		input.Body.LanguagePreference = "en"
		return input
	}()

	guardian, err := repo.CreateGuardian(ctx, guardianInput)
	if err != nil {
		t.Fatalf("Failed to create guardian: %v", err)
	}

	assert.NotNil(t, guardian)
	assert.Nil(t, err)
	assert.NotNil(t, guardian.ID)
	assert.NotNil(t, guardian.CreatedAt)
	assert.NotNil(t, guardian.UpdatedAt)
	assert.Equal(t, guardianInput.Body.Name, guardian.Name)

	// Verify we can retrieve the created guardian
	retrievedGuardian, err := repo.GetGuardianByID(ctx, guardian.ID)
	if err != nil {
		t.Fatalf("Failed to retrieve guardian: %v", err)
	}

	assert.NotNil(t, retrievedGuardian)
	assert.Equal(t, guardian.UserID, retrievedGuardian.UserID)
	assert.Equal(t, guardianInput.Body.Name, retrievedGuardian.Name)
}

func TestGuardianRepository_Create_Constraints(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()

	t.Run("Duplicate Email Failure", func(t *testing.T) {
		input1 := &models.CreateGuardianInput{}
		input1.Body.Name = "User One"
		input1.Body.Email = "duplicate@test.com"
		input1.Body.Username = "userone"

		_, err := repo.CreateGuardian(ctx, input1)
		assert.NoError(t, err)

		input2 := &models.CreateGuardianInput{}
		input2.Body.Name = "User Two"
		input2.Body.Email = "duplicate@test.com"
		input2.Body.Username = "usertwo"

		guardian, err := repo.CreateGuardian(ctx, input2)
		assert.Error(t, err)
		assert.Nil(t, guardian)
	})
}
