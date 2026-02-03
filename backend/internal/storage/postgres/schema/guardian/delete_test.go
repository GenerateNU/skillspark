package guardian

import (
	"context"
	"testing"

	"skillspark/internal/storage/postgres/testutil"

	"skillspark/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGuardianRepository_DeleteGuardian(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()
	t.Parallel()
	// delete a guardian that has a child that exists, should error
	guardian, err := repo.DeleteGuardian(ctx, uuid.MustParse("11111111-1111-1111-1111-111111111111"))

	assert.Error(t, err)
	assert.Nil(t, guardian)

	// create a guardian with no child
	guardian, err = repo.CreateGuardian(ctx,
		func() *models.CreateGuardianInput {
			input := &models.CreateGuardianInput{}
			input.Body.Name = "Delete Me"
			input.Body.Email = "deleteme@example.com"
			input.Body.Username = "delete"
			input.Body.LanguagePreference = "en"
			return input
		}())
	if err != nil {
		t.Fatalf("Failed to create guardian for deletion test: %v", err)
	}

	assert.NotNil(t, guardian)
	assert.NotNil(t, guardian.CreatedAt)
	assert.NotNil(t, guardian.UpdatedAt)

	// delete the guardian
	guardian, err = repo.DeleteGuardian(ctx, uuid.MustParse(guardian.ID.String()))
	if err != nil {
		t.Fatalf("Failed to delete guardian: %v", err)
	}

	assert.NotNil(t, guardian)
	assert.NotNil(t, guardian.CreatedAt)
	assert.NotNil(t, guardian.UpdatedAt)

	// verify the guardian was deleted
	guardian, err = repo.GetGuardianByID(ctx, guardian.ID)

	assert.Nil(t, guardian)
	assert.NotNil(t, err)
}

func TestGuardianRepository_Delete_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}
	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()

	guardian, err := repo.DeleteGuardian(ctx, uuid.New())
	assert.Error(t, err)
	assert.Nil(t, guardian)
	assert.Contains(t, err.Error(), "not found")
}
