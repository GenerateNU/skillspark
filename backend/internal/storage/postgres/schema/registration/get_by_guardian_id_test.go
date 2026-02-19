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

func TestGetRegistrationsByGuardianID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	r := CreateTestRegistration(t, ctx, testDB)

	input := &models.GetRegistrationsByGuardianIDInput{
		GuardianID: r.GuardianID,
	}

	result, err := repo.GetRegistrationsByGuardianID(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, result)
	assert.NotEmpty(t, result.Body.Registrations)

	for _, reg := range result.Body.Registrations {
		assert.Equal(t, r.GuardianID, reg.GuardianID)
		assert.NotEqual(t, uuid.Nil, reg.ID)
		assert.NotEqual(t, uuid.Nil, reg.ChildID)
		assert.NotEqual(t, uuid.Nil, reg.EventOccurrenceID)
		assert.NotEmpty(t, reg.Status)
		assert.NotEmpty(t, reg.EventName)
		assert.NotZero(t, reg.CreatedAt)
		assert.NotZero(t, reg.UpdatedAt)
		assert.NotZero(t, reg.OccurrenceStartTime)
		
		// Verify payment fields
		assert.NotEmpty(t, reg.StripePaymentIntentID)
		assert.NotEmpty(t, reg.StripeCustomerID)
		assert.NotEmpty(t, reg.OrgStripeAccountID)
		assert.NotEmpty(t, reg.StripePaymentMethodID)
		assert.NotZero(t, reg.TotalAmount)
		assert.NotZero(t, reg.ProviderAmount)
		assert.NotZero(t, reg.PlatformFeeAmount)
		assert.NotEmpty(t, reg.Currency)
		assert.NotEmpty(t, reg.PaymentIntentStatus)
	}
}

func TestGetRegistrationsByGuardianID_MultipleChildren(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardianID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	input := &models.GetRegistrationsByGuardianIDInput{
		GuardianID: guardianID,
	}

	result, err := repo.GetRegistrationsByGuardianID(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, result)
	assert.GreaterOrEqual(t, len(result.Body.Registrations), 2)

	childIDs := make(map[uuid.UUID]bool)
	for _, reg := range result.Body.Registrations {
		assert.Equal(t, &guardianID, reg.GuardianID)
		childIDs[reg.ChildID] = true
	}

	assert.GreaterOrEqual(t, len(childIDs), 1)
}

func TestGetRegistrationsByGuardianID_NoRegistrations(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardianID := uuid.New()

	input := &models.GetRegistrationsByGuardianIDInput{
		GuardianID: guardianID,
	}

	result, err := repo.GetRegistrationsByGuardianID(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, result)
	assert.Empty(t, result.Body.Registrations)
}

func TestGetRegistrationsByGuardianID_VerifyEventDetails(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	reg := CreateTestRegistration(t, ctx, testDB)
	guardianID := reg.GuardianID

	input := &models.GetRegistrationsByGuardianIDInput{
		GuardianID: guardianID,
	}

	result, err := repo.GetRegistrationsByGuardianID(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, result)
	assert.NotEmpty(t, result.Body.Registrations)

	for _, reg := range result.Body.Registrations {
		assert.NotEmpty(t, reg.EventName)
		assert.NotZero(t, reg.OccurrenceStartTime)
	}
}