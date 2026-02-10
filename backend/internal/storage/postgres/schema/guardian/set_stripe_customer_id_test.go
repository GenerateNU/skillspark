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

	require.Nil(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, testGuardian.ID, updated.ID)
	assert.NotNil(t, updated.StripeCustomerID)
	assert.Equal(t, stripeCustomerID, *updated.StripeCustomerID)

	fetched, getErr := repo.GetGuardianByID(ctx, testGuardian.ID)
	require.Nil(t, getErr)
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
	require.Nil(t, err)
	assert.Equal(t, firstCustomerID, *updated1.StripeCustomerID)

	secondCustomerID := "cus_second456"
	updated2, err := repo.SetStripeCustomerID(ctx, testGuardian.ID, secondCustomerID)
	require.Nil(t, err)
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

	require.NotNil(t, err)
	assert.Nil(t, updated)
}

func TestSetStripeCustomerID_DeletesPaymentMethods(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	guardianRepo := NewGuardianRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	testGuardian := CreateTestGuardian(t, ctx, testDB)
	firstCustomerID := "cus_first123"
	guardianRepo.SetStripeCustomerID(ctx, testGuardian.ID, firstCustomerID)

	_, err := testDB.Exec(ctx, `
		INSERT INTO guardian_payment_methods (guardian_id, stripe_payment_method_id, is_default)
		VALUES ($1, 'pm_old_card', true)
	`, testGuardian.ID)
	require.Nil(t, err)

	var countBefore int
	testDB.QueryRow(ctx, `SELECT COUNT(*) FROM guardian_payment_methods WHERE guardian_id = $1`, testGuardian.ID).Scan(&countBefore)
	assert.Equal(t, 1, countBefore)

	secondCustomerID := "cus_second456"
	updated, err := guardianRepo.SetStripeCustomerID(ctx, testGuardian.ID, secondCustomerID)
	require.Nil(t, err)
	assert.Equal(t, secondCustomerID, *updated.StripeCustomerID)

	var countAfter int
	testDB.QueryRow(ctx, `SELECT COUNT(*) FROM guardian_payment_methods WHERE guardian_id = $1`, testGuardian.ID).Scan(&countAfter)
	assert.Equal(t, 0, countAfter)
}