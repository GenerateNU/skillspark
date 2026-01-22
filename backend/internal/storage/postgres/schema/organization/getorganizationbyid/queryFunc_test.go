package getorganizationbyid

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExecute(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()

	org, err := Execute(ctx, testDB, uuid.MustParse("40000000-0000-0000-0000-000000000004"))

	require.Nil(t,err)
	assert.Equal(t, "Harmony Music School", org.Name)
	assert.True(t, org.Active)
}

func TestExecute_SecondOrganization(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()

	org, err := Execute(ctx, testDB, uuid.MustParse("40000000-0000-0000-0000-000000000003"))

	require.Nil(t, err)
	require.NotNil(t, org)
	assert.Equal(t, "Creative Arts Studio", org.Name)
	assert.True(t, org.Active)
}

func TestExecute_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()

	org, err := Execute(ctx, testDB, uuid.New())

	require.Error(t, err)
	assert.Nil(t, org)
}