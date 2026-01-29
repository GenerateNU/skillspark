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

	org, err := repo.GetOrganizationByID(ctx, uuid.MustParse("40000000-0000-0000-0000-000000000004"))

	require.Nil(t, err)
	assert.Equal(t, "Harmony Music School", org.Name)
	assert.True(t, org.Active)
}

func TestExecute_SecondOrganization(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	org, err := repo.GetOrganizationByID(ctx, uuid.MustParse("40000000-0000-0000-0000-000000000003"))

	require.Nil(t, err)
	require.NotNil(t, org)
	assert.Equal(t, "Creative Arts Studio", org.Name)
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
