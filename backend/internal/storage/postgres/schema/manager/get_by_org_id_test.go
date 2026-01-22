package manager

import (
	"context"
	"testing"

	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestManagerRepository_GetManagerByOrgID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewManagerRepository(testDB)
	ctx := context.Background()

	manager, err := repo.GetManagerByOrgID(ctx, uuid.MustParse("40000000-0000-0000-0000-000000000001"))

	assert.Nil(t, err)
	assert.NotNil(t, manager)
	assert.Equal(t, uuid.MustParse("c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f"), manager.UserID)
	assert.Equal(t, uuid.MustParse("40000000-0000-0000-0000-000000000001"), manager.OrganizationID)
	assert.Equal(t, "Director", manager.Role)

	managerTwo, err := repo.GetManagerByOrgID(ctx, uuid.MustParse("40000000-0000-0000-0000-000000000002"))

	assert.Nil(t, err)
	assert.NotNil(t, managerTwo)
	assert.Equal(t, uuid.MustParse("d0e1f2a3-b4c5-4d6e-7f8a-9b0c1d2e3f4a"), managerTwo.UserID)
	assert.Equal(t, uuid.MustParse("40000000-0000-0000-0000-000000000002"), managerTwo.OrganizationID)
	assert.Equal(t, "Head Coach", managerTwo.Role)

	managerThree, err := repo.GetManagerByOrgID(ctx, uuid.MustParse("00000000-0000-0000-0000-000000000000"))

	assert.NotNil(t, err)
	assert.Nil(t, managerThree)
}
