package guardian

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
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
		input.Body.UserID = uuid.MustParse("f2a3b4c5-d6e7-4f8a-9b0c-1d2e3f4a5b6c")
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
	assert.Equal(t, guardianInput.Body.UserID, guardian.UserID)

	// Verify we can retrieve the created guardian
	retrievedGuardian, err := repo.GetGuardianByID(ctx, guardian.ID)
	if err != nil {
		t.Fatalf("Failed to retrieve guardian: %v", err)
	}

	assert.NotNil(t, retrievedGuardian)
	assert.Equal(t, guardianInput.Body.UserID, retrievedGuardian.UserID)
}
