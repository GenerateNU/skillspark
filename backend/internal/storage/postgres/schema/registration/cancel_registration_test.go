package registration

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCancelRegistration(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)
	require.Equal(t, models.RegistrationStatusRegistered, created.Status)
	require.Nil(t, created.CancelledAt)

	// get enrolled count before cancellation
	var enrolledBefore int
	row := testDB.QueryRow(ctx, "SELECT curr_enrolled FROM event_occurrence WHERE id = $1", created.EventOccurrenceID)
	require.NoError(t, row.Scan(&enrolledBefore))

	input := &models.CancelRegistrationInput{
		ID: created.ID,
	}

	cancelled, err := repo.CancelRegistration(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, cancelled)
	assert.Equal(t, created.ID, cancelled.Body.Registration.ID)
	assert.Equal(t, models.RegistrationStatusCancelled, cancelled.Body.Registration.Status)
	assert.NotNil(t, cancelled.Body.Registration.CancelledAt)
	assert.Equal(t, created.ChildID, cancelled.Body.Registration.ChildID)
	assert.Equal(t, created.GuardianID, cancelled.Body.Registration.GuardianID)
	assert.Equal(t, created.EventOccurrenceID, cancelled.Body.Registration.EventOccurrenceID)
	assert.Equal(t, created.StripePaymentIntentID, cancelled.Body.Registration.StripePaymentIntentID)
	assert.Equal(t, created.TotalAmount, cancelled.Body.Registration.TotalAmount)
	assert.Equal(t, created.Currency, cancelled.Body.Registration.Currency)
	assert.NotEmpty(t, cancelled.Body.Registration.EventName)
	assert.NotZero(t, cancelled.Body.Registration.OccurrenceStartTime)

	// verify enrolled decremented
	var enrolledAfter int
	row = testDB.QueryRow(ctx, "SELECT curr_enrolled FROM event_occurrence WHERE id = $1", created.EventOccurrenceID)
	require.NoError(t, row.Scan(&enrolledAfter))
	assert.Equal(t, enrolledBefore-1, enrolledAfter)
}

func TestCancelRegistration_AlreadyCancelled(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)

	input := &models.CancelRegistrationInput{
		ID: created.ID,
	}
	_, err := repo.CancelRegistration(ctx, input)
	require.NoError(t, err)

	// second cancel should still succeed at the repo level
	cancelled, err := repo.CancelRegistration(ctx, input)
	require.NoError(t, err)
	require.NotNil(t, cancelled)
	assert.Equal(t, models.RegistrationStatusCancelled, cancelled.Body.Registration.Status)
	assert.NotNil(t, cancelled.Body.Registration.CancelledAt)
}

func TestCancelRegistration_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	input := &models.CancelRegistrationInput{
		ID: uuid.New(),
	}

	cancelled, err := repo.CancelRegistration(ctx, input)

	require.Error(t, err)
	assert.Nil(t, cancelled)
}

func TestCancelRegistration_VerifyPersistence(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)

	input := &models.CancelRegistrationInput{
		ID: created.ID,
	}
	cancelled, err := repo.CancelRegistration(ctx, input)
	require.NoError(t, err)

	fetched, err := repo.GetRegistrationByID(ctx, &models.GetRegistrationByIDInput{ID: created.ID}, nil)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, models.RegistrationStatusCancelled, fetched.Body.Status)
	assert.NotNil(t, fetched.Body.CancelledAt)
	assert.Equal(t, cancelled.Body.Registration.CancelledAt, fetched.Body.CancelledAt)
}

func TestCancelRegistration_WithPaymentIntentStatus(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)

	cancelledStatus := models.RegistrationStatusCancelled
	piStatus := "canceled"
	input := &models.CancelRegistrationInput{
		ID:                  created.ID,
		Status:              &cancelledStatus,
		PaymentIntentStatus: &piStatus,
	}

	cancelled, err := repo.CancelRegistration(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, cancelled)
	assert.Equal(t, models.RegistrationStatusCancelled, cancelled.Body.Registration.Status)
	assert.Equal(t, "canceled", cancelled.Body.Registration.PaymentIntentStatus)
	assert.NotNil(t, cancelled.Body.Registration.CancelledAt)
}