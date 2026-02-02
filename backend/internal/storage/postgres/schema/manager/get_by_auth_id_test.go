package manager

import (
	"context"
	"testing"

	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestManagerRepository_GetManagerByAuthID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewManagerRepository(testDB)
	ctx := context.Background()

	// Existing manager and user from seed (07_manager.sql and 01_profiles.sql)
	managerID := uuid.MustParse("50000000-0000-0000-0000-000000000001")
	userID := uuid.MustParse("c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f")

	// Generate a random AuthID
	authID := uuid.New()

	// Update the user to have this AuthID
	_, err := testDB.Exec(ctx, `UPDATE "user" SET auth_id = $1 WHERE id = $2`, authID, userID)
	if err != nil {
		t.Fatalf("Failed to update user auth_id: %v", err)
	}

	// Test Success Case
	manager, err := repo.GetManagerByAuthID(ctx, authID.String())
	assert.NoError(t, err)
	assert.NotNil(t, manager)
	assert.Equal(t, managerID, manager.ID)
	assert.Equal(t, userID, manager.UserID)
	assert.Equal(t, "Director", manager.Role)
	assert.False(t, manager.CreatedAt.IsZero())
	assert.False(t, manager.UpdatedAt.IsZero())

	// Test Not Found Case
	nonExistentAuthID := uuid.New().String()
	manager, err = repo.GetManagerByAuthID(ctx, nonExistentAuthID)
	assert.Error(t, err)
	assert.Nil(t, manager)
}
