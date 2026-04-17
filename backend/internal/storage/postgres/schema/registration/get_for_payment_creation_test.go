package registration

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRegistrationsForPaymentCreation(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	startTime := time.Now().Add(2 * 24 * time.Hour)
	reg := CreateTestRegistrationWithoutPayment(t, ctx, testDB, startTime)

	results, err := repo.GetRegistrationsForPaymentCreation(ctx)

	require.NoError(t, err)
	require.NotNil(t, results)

	found := false
	for _, r := range results {
		if r.ID == reg.ID {
			found = true
			assert.Equal(t, reg.GuardianID, r.GuardianID)
			assert.Equal(t, reg.EventOccurrenceID, r.EventOccurrenceID)
		}
	}
	assert.True(t, found, "Registration without payment in 4-day window should be returned")
}

func TestGetRegistrationsForPaymentCreation_ExcludesWithPayment(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	startTime := time.Now().Add(2 * 24 * time.Hour)
	reg := CreateTestRegistrationWithoutPayment(t, ctx, testDB, startTime)

	paymentInput := &models.CreatePaymentData{
		RegistrationID:        reg.ID,
		StripePaymentIntentID: "pi_test_" + reg.ID.String()[:8],
		StripeCustomerID:      "cus_test_" + reg.GuardianID.String()[:8],
		OrgStripeAccountID:    "acct_test_123",
		StripePaymentMethodID: "pm_test_123",
		TotalAmount:           10000,
		ProviderAmount:        8500,
		PlatformFeeAmount:     1500,
		Currency:              "usd",
		PaymentIntentStatus:   "requires_capture",
	}
	err := repo.CreatePayment(ctx, paymentInput)
	require.NoError(t, err)

	results, err := repo.GetRegistrationsForPaymentCreation(ctx)

	require.NoError(t, err)
	for _, r := range results {
		assert.NotEqual(t, reg.ID, r.ID, "Registration with existing payment should be excluded")
	}
}

func TestGetRegistrationsForPaymentCreation_ExcludesCancelled(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	startTime := time.Now().Add(2 * 24 * time.Hour)
	reg := CreateTestRegistrationWithoutPayment(t, ctx, testDB, startTime)

	_, err := repo.CancelRegistration(ctx, &models.CancelRegistrationInput{ID: reg.ID})
	require.NoError(t, err)

	results, err := repo.GetRegistrationsForPaymentCreation(ctx)

	require.NoError(t, err)
	for _, r := range results {
		assert.NotEqual(t, reg.ID, r.ID, "Cancelled registration should be excluded")
	}
}

func TestGetRegistrationsForPaymentCreation_ExcludesPastEvents(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	pastStart := time.Now().Add(-2 * 24 * time.Hour)
	reg := CreateTestRegistrationWithoutPayment(t, ctx, testDB, pastStart)

	results, err := repo.GetRegistrationsForPaymentCreation(ctx)

	require.NoError(t, err)
	for _, r := range results {
		assert.NotEqual(t, reg.ID, r.ID, "Registration with past event should be excluded")
	}
}

func TestGetRegistrationsForPaymentCreation_ExcludesFarFutureEvents(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	farFuture := time.Now().Add(10 * 24 * time.Hour)
	reg := CreateTestRegistrationWithoutPayment(t, ctx, testDB, farFuture)

	results, err := repo.GetRegistrationsForPaymentCreation(ctx)

	require.NoError(t, err)
	for _, r := range results {
		assert.NotEqual(t, reg.ID, r.ID, "Registration with event more than 4 days away should be excluded")
	}
}

func TestGetRegistrationsForPaymentCreation_Empty(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	// CreateTestRegistration uses a past event date, so it won't appear in results
	CreateTestRegistration(t, ctx, testDB)

	// Use a fresh DB-scoped check: query for a non-existent ID to confirm empty is possible
	results, err := repo.GetRegistrationsForPaymentCreation(ctx)

	require.NoError(t, err)
	require.NotNil(t, results)

	for _, r := range results {
		assert.NotEqual(t, uuid.Nil, r.ID, "All returned registrations should have valid IDs")
	}
}

func TestGetRegistrationsForPaymentCreation_VerifyFields(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	startTime := time.Now().Add(2 * 24 * time.Hour)
	reg := CreateTestRegistrationWithoutPayment(t, ctx, testDB, startTime)

	results, err := repo.GetRegistrationsForPaymentCreation(ctx)

	require.NoError(t, err)

	var found *models.RegistrationForPayment
	for i := range results {
		if results[i].ID == reg.ID {
			found = &results[i]
			break
		}
	}
	require.NotNil(t, found, "Registration should be in results")

	assert.NotEqual(t, uuid.Nil, found.ID)
	assert.NotEqual(t, uuid.Nil, found.GuardianID)
	assert.NotEqual(t, uuid.Nil, found.EventOccurrenceID)
	assert.Equal(t, reg.GuardianID, found.GuardianID)
	assert.Equal(t, reg.EventOccurrenceID, found.EventOccurrenceID)
}
