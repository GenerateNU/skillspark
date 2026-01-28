package guardian

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGuardianRepository_Update_David_Kim(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()

	guardianInput := func() *models.UpdateGuardianInput {
		input := &models.UpdateGuardianInput{}
		input.ID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
		input.Body.Name = "Updated David"
		input.Body.Email = "updated.david@example.com"
		input.Body.Username = "udavid"
		input.Body.LanguagePreference = "en"
		return input
	}()

	guardian, err := repo.UpdateGuardian(ctx, guardianInput)
	if err != nil {
		t.Fatalf("Failed to update guardian: %v", err)
	}

	assert.Nil(t, err)
	assert.NotNil(t, guardian)
	assert.Equal(t, guardianInput.Body.Name, guardian.Name)
	assert.NotNil(t, guardian.CreatedAt)
	assert.NotNil(t, guardian.UpdatedAt)
	assert.Equal(t, guardianInput.ID, guardian.ID)

	// Verify we can retrieve the updated guardian
	retrievedGuardian, err := repo.GetGuardianByID(ctx, guardianInput.ID)
	if err != nil {
		t.Fatalf("Failed to retrieve guardian: %v", err)
	}

	assert.NotNil(t, retrievedGuardian)
	assert.Equal(t, guardianInput.Body.Name, retrievedGuardian.Name)
	assert.Equal(t, guardianInput.ID, retrievedGuardian.ID)
}
