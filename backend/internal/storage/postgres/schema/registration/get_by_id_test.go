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

func TestGetRegistrationByID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	reg := CreateTestRegistration(t, ctx, testDB)
	registrationID := reg.ID

	getInput := &models.GetRegistrationByIDInput{
		ID: registrationID,
	}

	retrieved, err := repo.GetRegistrationByID(ctx, getInput)

	require.Nil(t, err)
	require.NotNil(t, retrieved)
	assert.Equal(t, registrationID, retrieved.Body.ID)
	assert.NotEqual(t, uuid.Nil, retrieved.Body.ChildID)
	assert.NotEqual(t, uuid.Nil, retrieved.Body.GuardianID)
	assert.NotEqual(t, uuid.Nil, retrieved.Body.EventOccurrenceID)
	assert.NotEmpty(t, retrieved.Body.Status)
	assert.NotEmpty(t, retrieved.Body.EventName)
	assert.NotZero(t, retrieved.Body.CreatedAt)
	assert.NotZero(t, retrieved.Body.UpdatedAt)
	assert.NotZero(t, retrieved.Body.OccurrenceStartTime)
	
	// Verify payment fields
	assert.NotEmpty(t, retrieved.Body.StripePaymentIntentID)
	assert.NotEmpty(t, retrieved.Body.StripeCustomerID)
	assert.NotEmpty(t, retrieved.Body.OrgStripeAccountID)
	assert.NotEmpty(t, retrieved.Body.StripePaymentMethodID)
	assert.NotZero(t, retrieved.Body.TotalAmount)
	assert.NotZero(t, retrieved.Body.ProviderAmount)
	assert.NotZero(t, retrieved.Body.PlatformFeeAmount)
	assert.NotEmpty(t, retrieved.Body.Currency)
	assert.NotEmpty(t, retrieved.Body.PaymentIntentStatus)
	assert.Nil(t, retrieved.Body.CancelledAt)
	assert.Nil(t, retrieved.Body.PaidAt)
}

func TestGetRegistrationByID_VerifyEventDetails(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	reg := CreateTestRegistration(t, ctx, testDB)

	registrationID := reg.ID

	getInput := &models.GetRegistrationByIDInput{
		ID: registrationID,
	}

	retrieved, err := repo.GetRegistrationByID(ctx, getInput)

	require.Nil(t, err)
	require.NotNil(t, retrieved)
	assert.NotEmpty(t, retrieved.Body.EventName)
	assert.NotZero(t, retrieved.Body.OccurrenceStartTime)
}

func TestGetRegistrationByID_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	getInput := &models.GetRegistrationByIDInput{
		ID: uuid.New(),
	}

	retrieved, err := repo.GetRegistrationByID(ctx, getInput)

	require.NotNil(t, err)
	require.Nil(t, retrieved)
}

func TestGetRegistrationByID_MultipleDifferentRegistrations(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	reg1 := CreateTestRegistration(t, ctx, testDB)
	reg2 := CreateTestRegistration(t, ctx, testDB)

	getInput1 := &models.GetRegistrationByIDInput{
		ID: reg1.ID,
	}
	retrieved1, err := repo.GetRegistrationByID(ctx, getInput1)
	require.Nil(t, err)
	require.NotNil(t, retrieved1)

	getInput2 := &models.GetRegistrationByIDInput{
		ID: reg2.ID,
	}
	retrieved2, err := repo.GetRegistrationByID(ctx, getInput2)
	require.Nil(t, err)
	require.NotNil(t, retrieved2)

	assert.NotEqual(t, retrieved1.Body.ID, retrieved2.Body.ID)
}