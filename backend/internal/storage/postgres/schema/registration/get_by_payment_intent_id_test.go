package registration

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRegistrationByPaymentIntentID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)
	require.NotEmpty(t, created.StripePaymentIntentID)

	registration, err := repo.GetRegistrationByPaymentIntentID(ctx, created.StripePaymentIntentID)

	require.NoError(t, err)
	require.NotNil(t, registration)
	assert.Equal(t, created.ID, registration.ID)
	assert.Equal(t, created.ChildID, registration.ChildID)
	assert.Equal(t, created.GuardianID, registration.GuardianID)
	assert.Equal(t, created.EventOccurrenceID, registration.EventOccurrenceID)
	assert.Equal(t, created.StripePaymentIntentID, registration.StripePaymentIntentID)
	assert.Equal(t, created.StripeCustomerID, registration.StripeCustomerID)
	assert.Equal(t, created.OrgStripeAccountID, registration.OrgStripeAccountID)
	assert.Equal(t, created.TotalAmount, registration.TotalAmount)
	assert.Equal(t, created.Currency, registration.Currency)
	assert.Equal(t, created.Status, registration.Status)
	assert.NotEmpty(t, registration.EventName)
	assert.NotZero(t, registration.OccurrenceStartTime)
}

func TestGetRegistrationByPaymentIntentID_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	registration, err := repo.GetRegistrationByPaymentIntentID(ctx, "pi_doesnotexist")

	require.Error(t, err)
	assert.Nil(t, registration)
}

func TestGetRegistrationByPaymentIntentID_EmptyID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	registration, err := repo.GetRegistrationByPaymentIntentID(ctx, "")

	require.Error(t, err)
	assert.Nil(t, registration)
}

func TestGetRegistrationByPaymentIntentID_AfterCancel(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)

	_, err := repo.CancelRegistration(ctx, &models.CancelRegistrationInput{ID: created.ID})
	require.NoError(t, err)

	registration, err := repo.GetRegistrationByPaymentIntentID(ctx, created.StripePaymentIntentID)

	require.NoError(t, err)
	require.NotNil(t, registration)
	assert.Equal(t, created.ID, registration.ID)
	assert.Equal(t, models.RegistrationStatusCancelled, registration.Status)
	assert.NotNil(t, registration.CancelledAt)
}
