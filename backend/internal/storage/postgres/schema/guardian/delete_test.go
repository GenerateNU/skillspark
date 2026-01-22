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
	// delete a guardian that has a child that exists, should error
	guardian, err := repo.DeleteGuardian(ctx, uuid.MustParse("11111111-1111-1111-1111-111111111111"))

	assert.Error(t, err)
	assert.Nil(t, guardian)

	// create a guardian with no child
	guardian, err = repo.CreateGuardian(ctx,
		func() *models.CreateGuardianInput {
			input := &models.CreateGuardianInput{}
			input.Body.UserID = uuid.MustParse("c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f")
			return input
		}())
	if err != nil {
		t.Fatalf("Failed to create guardian for deletion test: %v", err)
	}

	assert.NotNil(t, guardian)
	assert.Equal(t, uuid.MustParse("c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f"), guardian.UserID)
	assert.NotNil(t, guardian.CreatedAt)
	assert.NotNil(t, guardian.UpdatedAt)

	// delete the guardian
	guardian, err = repo.DeleteGuardian(ctx, uuid.MustParse(guardian.ID.String()))
	if err != nil {
		t.Fatalf("Failed to delete guardian: %v", err)
	}

	assert.NotNil(t, guardian)
	assert.Equal(t, uuid.MustParse("c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f"), guardian.UserID)
	assert.NotNil(t, guardian.CreatedAt)
	assert.NotNil(t, guardian.UpdatedAt)

	// verify the guardian was deleted
	guardian, err = repo.GetGuardianByID(ctx, uuid.MustParse("c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f"))

	assert.Nil(t, guardian)
	assert.NotNil(t, err)

}
