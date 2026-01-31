package guardian

import (
	"context"
	"testing"

	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGuardianRepository_GetGuardianByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	// get a guardian that exists
	guardian, err := repo.GetGuardianByID(ctx, uuid.MustParse("11111111-1111-1111-1111-111111111111"))
	if err != nil {
		t.Fatalf("Failed to get guardian by id: %v", err)
	}

	assert.NotNil(t, guardian)
	assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111111"), guardian.ID)
	assert.NotNil(t, guardian.CreatedAt)
	assert.NotNil(t, guardian.UpdatedAt)

	// get a guardian that does not exist
	guardian, err = repo.GetGuardianByID(ctx, uuid.MustParse("22222222-2222-2222-2222-222222222223"))
	assert.Nil(t, guardian)
	assert.NotNil(t, err)
}
