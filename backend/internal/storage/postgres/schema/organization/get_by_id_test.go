package organization

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetById(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	testorg := CreateTestOrganization(t, ctx, testDB)

	org, err := repo.GetOrganizationByID(ctx, testorg.ID)

	require.Nil(t, err)
	assert.Equal(t, "Test Corp", org.Name)
	assert.True(t, org.Active)
}

func TestExecute_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	org, err := repo.GetOrganizationByID(ctx, uuid.New())

	require.Error(t, err)
	assert.Nil(t, org)
}
