package guardian

import (
	"context"
	"testing"

	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGuardianRepository_GetGuardianByChildID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()

	// get a guardian by a child id
	guardianFromChildID, err := repo.GetGuardianByChildID(ctx, uuid.MustParse("30000000-0000-0000-0000-000000000009"))
	if err != nil {
		t.Fatalf("Failed to get guardian by child id: %v", err)
	}

	assert.NotNil(t, guardianFromChildID)
	assert.Nil(t, err)
	assert.NotNil(t, guardianFromChildID.CreatedAt)
	assert.NotNil(t, guardianFromChildID.UpdatedAt)

	// pull guardian by the actual id
	guardian, err := repo.GetGuardianByID(ctx, uuid.MustParse("66666666-6666-6666-6666-666666666666"))
	if err != nil {
		t.Fatalf("Failed to get guardian by id: %v", err)
	}

	// compare the two guardians
	assert.Equal(t, guardianFromChildID.ID, guardian.ID)
	assert.Equal(t, guardianFromChildID.UserID, guardian.UserID)
}

func TestGuardianRepository_GetGuardianByChildID_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()

	randomChildID := uuid.New()

	guardian, err := repo.GetGuardianByChildID(ctx, randomChildID)

	assert.Error(t, err)
	assert.Nil(t, guardian)
	assert.Contains(t, err.Error(), "not found")
}
