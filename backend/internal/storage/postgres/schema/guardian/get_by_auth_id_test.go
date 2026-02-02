package guardian

import (
	"context"
	"testing"

	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGuardianRepository_GetGuardianByAuthID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()

	// Existing guardian and user from seed (03_guardian.sql and 01_profiles.sql)
	guardianID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	userID := uuid.MustParse("a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d")

	// Generate a random AuthID
	authID := uuid.New()

	// Update the user to have this AuthID
	_, err := testDB.Exec(ctx, `UPDATE "user" SET auth_id = $1 WHERE id = $2`, authID, userID)
	if err != nil {
		t.Fatalf("Failed to update user auth_id: %v", err)
	}

	// Test Success Case
	guardian, err := repo.GetGuardianByAuthID(ctx, authID.String())
	assert.NoError(t, err)
	assert.NotNil(t, guardian)
	assert.Equal(t, guardianID, guardian.ID)
	assert.Equal(t, userID, guardian.UserID)
	assert.False(t, guardian.CreatedAt.IsZero())
	assert.False(t, guardian.UpdatedAt.IsZero())

	// Test Not Found Case
	nonExistentAuthID := uuid.New().String()
	guardian, err = repo.GetGuardianByAuthID(ctx, nonExistentAuthID)
	assert.Error(t, err)
	assert.Nil(t, guardian)
	// Check if error is NotFound type if possible, or just check it's not nil
	// Based on implementation it returns errs.NotFound which might satisfy checks differently,
	// but asserting Error is a good start.
}
