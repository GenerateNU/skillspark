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

	// delete a guardian that has children that exist, should delete guardian and children
	guardian, err := repo.GetGuardianByChildID(ctx, uuid.MustParse("30000000-0000-0000-0000-000000000001")) // child 1
	assert.Nil(t, err)
	assert.NotNil(t, guardian)
	assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111111"), guardian.ID)

	guardian, err = repo.GetGuardianByChildID(ctx, uuid.MustParse("30000000-0000-0000-0000-000000000002")) // child 2
	assert.Nil(t, err)
	assert.NotNil(t, guardian)
	assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111111"), guardian.ID)

	guardian, err = repo.DeleteGuardian(ctx, uuid.MustParse("11111111-1111-1111-1111-111111111111"), nil)
	assert.Nil(t, err)
	assert.NotNil(t, guardian)

	guardian, err = repo.GetGuardianByID(ctx, guardian.ID) // guardian deleted
	assert.Nil(t, guardian)
	assert.NotNil(t, err)

	guardian, err = repo.GetGuardianByChildID(ctx, uuid.MustParse("30000000-0000-0000-0000-000000000001")) // child 1 deleted
	assert.Nil(t, guardian)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Child with id: 30000000-0000-0000-0000-000000000001 not found")

	guardian, err = repo.GetGuardianByChildID(ctx, uuid.MustParse("30000000-0000-0000-0000-000000000002")) // child 2 deleted
	assert.Nil(t, guardian)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Child with id: 30000000-0000-0000-0000-000000000002 not found")

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
	guardian, err = repo.DeleteGuardian(ctx, uuid.MustParse(guardian.ID.String()), nil)
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

	guardian, err := repo.DeleteGuardian(ctx, uuid.New(), nil)
	assert.Error(t, err)
	assert.Nil(t, guardian)
}
