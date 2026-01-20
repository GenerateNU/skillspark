package guardian

import (
	"context"
	"testing"

	"skillspark/internal/storage/postgres/testutil"

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
	// delete a guardian that exists
	guardian, err := repo.DeleteGuardian(ctx, uuid.MustParse("11111111-1111-1111-1111-111111111111"))

	// TODO: fix cascade to delete child as well 
	if err != nil {
		t.Fatalf("Failed to delete guardian: %v", err)
	}

	assert.NotNil(t, guardian)
	assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111111"), guardian.ID)
	assert.Equal(t, uuid.MustParse("a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d"), guardian.UserID)
	assert.NotNil(t, guardian.CreatedAt)
	assert.NotNil(t, guardian.UpdatedAt)

	// verify the guardian was deleted
	guardian, err = repo.GetGuardianByID(ctx, uuid.MustParse("11111111-1111-1111-1111-111111111111"))

	assert.Nil(t, guardian)
	assert.NotNil(t, err)
}