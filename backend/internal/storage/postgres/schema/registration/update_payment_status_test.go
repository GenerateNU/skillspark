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

func TestUpdateRegistrationPaymentStatus(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)
	require.Equal(t, "requires_capture", created.PaymentIntentStatus)
	require.Nil(t, created.PaidAt)

	input := &models.UpdateRegistrationPaymentStatusInput{
		ID: created.ID,
	}
	input.Body.PaymentIntentStatus = "succeeded"

	updated, err := repo.UpdateRegistrationPaymentStatus(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, created.ID, updated.Body.ID)
	assert.Equal(t, "succeeded", updated.Body.PaymentIntentStatus)
	assert.NotNil(t, updated.Body.PaidAt)
	assert.NotZero(t, updated.Body.PaidAt)
	
	assert.Equal(t, created.ChildID, updated.Body.ChildID)
	assert.Equal(t, created.GuardianID, updated.Body.GuardianID)
	assert.Equal(t, created.EventOccurrenceID, updated.Body.EventOccurrenceID)
	assert.Equal(t, created.Status, updated.Body.Status)
	assert.Equal(t, created.StripePaymentIntentID, updated.Body.StripePaymentIntentID)
	assert.Equal(t, created.TotalAmount, updated.Body.TotalAmount)
	assert.Equal(t, created.Currency, updated.Body.Currency)
	assert.NotEmpty(t, updated.Body.EventName)
	assert.NotZero(t, updated.Body.OccurrenceStartTime)
}

func TestUpdateRegistrationPaymentStatus_RequiresCapture(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)

	input := &models.UpdateRegistrationPaymentStatusInput{
		ID: created.ID,
	}
	input.Body.PaymentIntentStatus = "requires_capture"

	updated, err := repo.UpdateRegistrationPaymentStatus(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, "requires_capture", updated.Body.PaymentIntentStatus)
	assert.Nil(t, updated.Body.PaidAt)
}

func TestUpdateRegistrationPaymentStatus_Processing(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)

	input := &models.UpdateRegistrationPaymentStatusInput{
		ID: created.ID,
	}
	input.Body.PaymentIntentStatus = "processing"

	updated, err := repo.UpdateRegistrationPaymentStatus(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, "processing", updated.Body.PaymentIntentStatus)
	assert.Nil(t, updated.Body.PaidAt)
}

func TestUpdateRegistrationPaymentStatus_Canceled(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)

	input := &models.UpdateRegistrationPaymentStatusInput{
		ID: created.ID,
	}
	input.Body.PaymentIntentStatus = "canceled"

	updated, err := repo.UpdateRegistrationPaymentStatus(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, "canceled", updated.Body.PaymentIntentStatus)
	assert.Nil(t, updated.Body.PaidAt)
}

func TestUpdateRegistrationPaymentStatus_RequiresPaymentMethod(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)

	input := &models.UpdateRegistrationPaymentStatusInput{
		ID: created.ID,
	}
	input.Body.PaymentIntentStatus = "requires_payment_method"

	updated, err := repo.UpdateRegistrationPaymentStatus(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, "requires_payment_method", updated.Body.PaymentIntentStatus)
	assert.Nil(t, updated.Body.PaidAt)
}

func TestUpdateRegistrationPaymentStatus_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	nonExistentID := uuid.New()

	input := &models.UpdateRegistrationPaymentStatusInput{
		ID: nonExistentID,
	}
	input.Body.PaymentIntentStatus = "succeeded"

	updated, err := repo.UpdateRegistrationPaymentStatus(ctx, input)

	require.NotNil(t, err)
	assert.Nil(t, updated)
}

func TestUpdateRegistrationPaymentStatus_VerifyPersistence(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)

	input := &models.UpdateRegistrationPaymentStatusInput{
		ID: created.ID,
	}
	input.Body.PaymentIntentStatus = "succeeded"

	updated, err := repo.UpdateRegistrationPaymentStatus(ctx, input)
	require.Nil(t, err)

	getInput := &models.GetRegistrationByIDInput{
		ID: created.ID,
	}
	fetched, err := repo.GetRegistrationByID(ctx, getInput)

	require.Nil(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, "succeeded", fetched.Body.PaymentIntentStatus)
	assert.NotNil(t, fetched.Body.PaidAt)
	assert.Equal(t, updated.Body.PaidAt, fetched.Body.PaidAt)
}

func TestUpdateRegistrationPaymentStatus_MultipleUpdates(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)

	input1 := &models.UpdateRegistrationPaymentStatusInput{
		ID: created.ID,
	}
	input1.Body.PaymentIntentStatus = "processing"

	updated1, err := repo.UpdateRegistrationPaymentStatus(ctx, input1)
	require.Nil(t, err)
	assert.Equal(t, "processing", updated1.Body.PaymentIntentStatus)
	assert.Nil(t, updated1.Body.PaidAt)

	input2 := &models.UpdateRegistrationPaymentStatusInput{
		ID: created.ID,
	}
	input2.Body.PaymentIntentStatus = "succeeded"

	updated2, err := repo.UpdateRegistrationPaymentStatus(ctx, input2)
	require.Nil(t, err)
	assert.Equal(t, "succeeded", updated2.Body.PaymentIntentStatus)
	assert.NotNil(t, updated2.Body.PaidAt)
}

func TestUpdateRegistrationPaymentStatus_PaidAtOnlySetOnSuccess(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)

	statuses := []string{"processing", "requires_payment_method", "requires_capture", "canceled"}

	for _, status := range statuses {
		input := &models.UpdateRegistrationPaymentStatusInput{
			ID: created.ID,
		}
		input.Body.PaymentIntentStatus = status

		updated, err := repo.UpdateRegistrationPaymentStatus(ctx, input)
		require.Nil(t, err)
		assert.Equal(t, status, updated.Body.PaymentIntentStatus)
		assert.Nil(t, updated.Body.PaidAt, "paid_at should be nil for status: %s", status)
	}

	input := &models.UpdateRegistrationPaymentStatusInput{
		ID: created.ID,
	}
	input.Body.PaymentIntentStatus = "succeeded"

	updated, err := repo.UpdateRegistrationPaymentStatus(ctx, input)
	require.Nil(t, err)
	assert.Equal(t, "succeeded", updated.Body.PaymentIntentStatus)
	assert.NotNil(t, updated.Body.PaidAt)
}

func TestUpdateRegistrationPaymentStatus_DoesNotAffectCancellation(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)
	cancelInput := &models.CancelRegistrationInput{
		ID: created.ID,
	}
	_, err := repo.CancelRegistration(ctx, cancelInput)
	require.Nil(t, err)

	input := &models.UpdateRegistrationPaymentStatusInput{
		ID: created.ID,
	}
	input.Body.PaymentIntentStatus = "succeeded"

	updated, err := repo.UpdateRegistrationPaymentStatus(ctx, input)
	require.Nil(t, err)
	assert.Equal(t, models.RegistrationStatusCancelled, updated.Body.Status)
	assert.NotNil(t, updated.Body.CancelledAt)
	assert.Equal(t, "succeeded", updated.Body.PaymentIntentStatus)
	assert.NotNil(t, updated.Body.PaidAt)
}