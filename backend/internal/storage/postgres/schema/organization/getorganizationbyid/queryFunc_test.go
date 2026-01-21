package getorganizationbyid

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestExecute(t *testing.T) {
// 	testDB := testutil.SetupTestDB(t)
// 	ctx := context.Background()

// 	org, err := Execute(ctx, testDB, uuid.MustParse("6b525ef8-eeea-4f26-8a66-61163cae5dc8"))

// 	require.Nil(t,err)
// 	assert.Equal(t, "Google", org.Name)
// 	assert.True(t, org.Active)
// }

// func TestExecute_SecondOrganization(t *testing.T) {
// 	testDB := testutil.SetupTestDB(t)
// 	ctx := context.Background()

// 	org, err := Execute(ctx, testDB, uuid.MustParse("40000000-0000-0000-0000-000000000003"))

// 	require.NoError(t, err)
// 	require.NotNil(t, org)
// 	assert.Equal(t, "Creative Arts Studio", org.Name)
// 	assert.True(t, org.Active)
// }

// func TestExecute_InactiveOrganization(t *testing.T) {
// 	testDB := testutil.SetupTestDB(t)
// 	ctx := context.Background()

// 	org, err := Execute(ctx, testDB, uuid.MustParse("2dcf6e5b-5c6e-4a98-8ac9-8b03edf0aed4"))

// 	require.NoError(t, err)
// 	require.NotNil(t, org)
// 	assert.Equal(t, "TESTING123", org.Name)
// 	assert.False(t, org.Active)  // This one is inactive
// }

// func TestExecute_AnotherInactiveOrg(t *testing.T) {
// 	testDB := testutil.SetupTestDB(t)
// 	ctx := context.Background()

// 	org, err := Execute(ctx, testDB, uuid.MustParse("40000000-0000-0000-0000-000000000007"))

// 	require.NoError(t, err)
// 	require.NotNil(t, org)
// 	assert.Equal(t, "Language Learning Center", org.Name)
// 	assert.False(t, org.Active)
// }

func TestExecute_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()

	org, err := Execute(ctx, testDB, uuid.New())

	require.Error(t, err)
	assert.Nil(t, org)
}