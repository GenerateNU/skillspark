package registration

import (
	"context"
	"skillspark/internal/models"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"

	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRegistrationsByEventOccurrenceID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	r := CreateTestRegistration(t, ctx, testDB)

	input := &models.GetRegistrationsByEventOccurrenceIDInput{
		EventOccurrenceID: r.EventOccurrenceID,
	}

	result, err := repo.GetRegistrationsByEventOccurrenceID(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, result)
	assert.NotEmpty(t, result.Body.Registrations)

	for _, reg := range result.Body.Registrations {
		assert.Equal(t, r.EventOccurrenceID, reg.EventOccurrenceID)
		assert.NotEqual(t, uuid.Nil, reg.ID)
		assert.NotEqual(t, uuid.Nil, reg.ChildID)
		assert.NotEqual(t, uuid.Nil, reg.GuardianID)
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

func TestGetRegistrationsByEventOccurrenceID_MultipleRegistrations(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	eventOccurrenceID := uuid.MustParse("70000000-0000-0000-0000-000000000001")

	input := &models.GetRegistrationsByEventOccurrenceIDInput{
		EventOccurrenceID: eventOccurrenceID,
	}

	result, err := repo.GetRegistrationsByEventOccurrenceID(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, result)
	assert.GreaterOrEqual(t, len(result.Body.Registrations), 2)

	registrationIDs := make(map[uuid.UUID]bool)
	for _, reg := range result.Body.Registrations {
		assert.Equal(t, eventOccurrenceID, reg.EventOccurrenceID)
		registrationIDs[reg.ID] = true
	}

	assert.Equal(t, len(result.Body.Registrations), len(registrationIDs))
}

func TestGetRegistrationsByEventOccurrenceID_NoRegistrations(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)

	input := &models.GetRegistrationsByEventOccurrenceIDInput{
		EventOccurrenceID: occurrence.ID,
	}

	result, err := repo.GetRegistrationsByEventOccurrenceID(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, result)
	assert.Empty(t, result.Body.Registrations)
}

func TestGetRegistrationsByEventOccurrenceID_VerifyEventDetails(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	r := CreateTestRegistration(t, ctx, testDB)

	input := &models.GetRegistrationsByEventOccurrenceIDInput{
		EventOccurrenceID: r.EventOccurrenceID,
	}

	result, err := repo.GetRegistrationsByEventOccurrenceID(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, result)
	assert.NotEmpty(t, result.Body.Registrations)

	for _, reg := range result.Body.Registrations {
		assert.NotEmpty(t, reg.EventName)
		assert.NotZero(t, reg.OccurrenceStartTime)
	}
}