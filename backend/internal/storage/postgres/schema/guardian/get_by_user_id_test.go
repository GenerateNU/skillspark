package guardian

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGuardianRepository_GetGuardianByUserID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	// get a guardian that exists
	guardian, err := repo.GetGuardianByUserID(ctx, uuid.MustParse("b8c9d0e1-f2a3-4b4c-5d6e-7f8a9b0c1d2e"))
	if err != nil {
		t.Fatalf("Failed to get guardian by id: %v", err)
	}

	assert.NotNil(t, guardian)
	assert.Equal(t, uuid.MustParse("88888888-8888-8888-8888-888888888888"), guardian.ID)
	assert.NotNil(t, guardian.CreatedAt)
	assert.NotNil(t, guardian.UpdatedAt)

	// get a guardian that does not exist
	guardian, err = repo.GetGuardianByID(ctx, uuid.MustParse("22222222-2222-2222-2222-222222222223"))
	assert.Nil(t, guardian)
	assert.NotNil(t, err)
}
