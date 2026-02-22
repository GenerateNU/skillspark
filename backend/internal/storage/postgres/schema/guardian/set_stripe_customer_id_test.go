package guardian

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetStripeCustomerID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	testGuardian := CreateTestGuardian(t, ctx, testDB)
	require.NotNil(t, testGuardian)
	require.Nil(t, testGuardian.StripeCustomerID)

	stripeCustomerID := "cus_test123abc"
	updated, err := repo.SetStripeCustomerID(ctx, testGuardian.ID, stripeCustomerID)

	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, testGuardian.ID, updated.ID)
	assert.NotNil(t, updated.StripeCustomerID)
	assert.Equal(t, stripeCustomerID, *updated.StripeCustomerID)

	fetched, err := repo.GetGuardianByID(ctx, testGuardian.ID)
	require.NoError(t, err)
	require.NotNil(t, fetched.StripeCustomerID)
	assert.Equal(t, stripeCustomerID, *fetched.StripeCustomerID)
}

func TestSetStripeCustomerID_UpdateExisting(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	testGuardian := CreateTestGuardian(t, ctx, testDB)

	firstCustomerID := "cus_first123"
	updated1, err := repo.SetStripeCustomerID(ctx, testGuardian.ID, firstCustomerID)
	require.NoError(t, err)
	require.NotNil(t, updated1.StripeCustomerID)
	assert.Equal(t, firstCustomerID, *updated1.StripeCustomerID)

	secondCustomerID := "cus_second456"
	updated2, err := repo.SetStripeCustomerID(ctx, testGuardian.ID, secondCustomerID)
	require.NoError(t, err)
	require.NotNil(t, updated2.StripeCustomerID)
	assert.Equal(t, secondCustomerID, *updated2.StripeCustomerID)
}

func TestSetStripeCustomerID_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	nonExistentID := uuid.New()
	stripeCustomerID := "cus_test123"

	updated, err := repo.SetStripeCustomerID(ctx, nonExistentID, stripeCustomerID)

	require.Error(t, err)
	assert.Nil(t, updated)
}