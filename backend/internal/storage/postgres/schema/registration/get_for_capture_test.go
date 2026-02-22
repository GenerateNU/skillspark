package registration

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRegistrationsForCapture(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	reg := CreateTestRegistration(t, ctx, testDB)
	require.Equal(t, "requires_capture", reg.PaymentIntentStatus)
	require.Equal(t, models.RegistrationStatusRegistered, reg.Status)

	startWindow := reg.OccurrenceStartTime.Add(-1 * time.Hour)
	endWindow := reg.OccurrenceStartTime.Add(1 * time.Hour)

	registrations, err := repo.GetRegistrationsForCapture(ctx, startWindow, endWindow)

	require.NoError(t, err)
	require.NotNil(t, registrations)
	
	found := false
	for _, r := range registrations {
		if r.ID == reg.ID {
			found = true
			assert.Equal(t, "requires_capture", r.PaymentIntentStatus)
			assert.Equal(t, models.RegistrationStatusRegistered, r.Status)
			assert.NotEmpty(t, r.StripePaymentIntentID)
			assert.NotEmpty(t, r.OrgStripeAccountID)
		}
	}
	assert.True(t, found, "Created registration should be in capture window")
}

func TestGetRegistrationsForCapture_ExcludesCaptured(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	reg := CreateTestRegistration(t, ctx, testDB)

	updateInput := &models.UpdateRegistrationPaymentStatusInput{
		ID: reg.ID,
	}
	updateInput.Body.PaymentIntentStatus = "succeeded"
	_, err := repo.UpdateRegistrationPaymentStatus(ctx, updateInput)
	require.NoError(t, err)

	startWindow := reg.OccurrenceStartTime.Add(-1 * time.Hour)
	endWindow := reg.OccurrenceStartTime.Add(1 * time.Hour)

	registrations, err := repo.GetRegistrationsForCapture(ctx, startWindow, endWindow)

	require.NoError(t, err)
	for _, r := range registrations {
		assert.NotEqual(t, reg.ID, r.ID, "Succeeded payment should not be in capture list")
	}
}

func TestGetRegistrationsForCapture_ExcludesCancelled(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	reg := CreateTestRegistration(t, ctx, testDB)

	cancelInput := &models.CancelRegistrationInput{
		ID: reg.ID,
	}
	_, err := repo.CancelRegistration(ctx, cancelInput)
	require.NoError(t, err)

	startWindow := reg.OccurrenceStartTime.Add(-1 * time.Hour)
	endWindow := reg.OccurrenceStartTime.Add(1 * time.Hour)

	registrations, err := repo.GetRegistrationsForCapture(ctx, startWindow, endWindow)

	require.NoError(t, err)
	for _, r := range registrations {
		assert.NotEqual(t, reg.ID, r.ID, "Cancelled registration should not be in capture list")
	}
}

func TestGetRegistrationsForCapture_TimeWindowFiltering(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	reg := CreateTestRegistration(t, ctx, testDB)

	tooEarly := reg.OccurrenceStartTime.Add(-10 * time.Hour)
	tooLate := reg.OccurrenceStartTime.Add(-8 * time.Hour)

	registrations, err := repo.GetRegistrationsForCapture(ctx, tooEarly, tooLate)

	require.NoError(t, err)
	for _, r := range registrations {
		assert.NotEqual(t, reg.ID, r.ID, "Registration outside time window should not be included")
	}
}

func TestGetRegistrationsForCapture_Empty(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	farFuture := time.Now().Add(365 * 24 * time.Hour)
	endWindow := farFuture.Add(1 * time.Hour)

	registrations, err := repo.GetRegistrationsForCapture(ctx, farFuture, endWindow)

	require.NoError(t, err)
	require.NotNil(t, registrations)
	assert.Equal(t, 0, len(registrations))
}

func TestGetRegistrationsForCapture_VerifyAllFields(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	reg := CreateTestRegistration(t, ctx, testDB)

	startWindow := reg.OccurrenceStartTime.Add(-1 * time.Hour)
	endWindow := reg.OccurrenceStartTime.Add(1 * time.Hour)

	registrations, err := repo.GetRegistrationsForCapture(ctx, startWindow, endWindow)

	require.NoError(t, err)
	require.NotEmpty(t, registrations)

	for _, r := range registrations {
		assert.NotEmpty(t, r.ID)
		assert.NotEmpty(t, r.StripePaymentIntentID)
		assert.NotEmpty(t, r.OrgStripeAccountID)
		assert.NotEmpty(t, r.EventName)
		assert.NotZero(t, r.OccurrenceStartTime)
		assert.NotZero(t, r.TotalAmount)
		assert.Equal(t, "requires_capture", r.PaymentIntentStatus)
		assert.Equal(t, models.RegistrationStatusRegistered, r.Status)
	}
}