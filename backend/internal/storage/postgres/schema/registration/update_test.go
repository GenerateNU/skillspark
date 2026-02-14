package registration

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/child"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"
	"skillspark/internal/storage/postgres/schema/guardian"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateRegistration(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	childID := child.CreateTestChild(t, ctx, testDB).ID
	newChildID := child.CreateTestChild(t, ctx, testDB).ID
	guardianID := guardian.CreateTestGuardian(t, ctx, testDB).ID
	occurrenceID := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB).ID

	createInput := &models.CreateRegistrationWithPaymentData{
		ChildID:                childID,
		GuardianID:             guardianID,
		EventOccurrenceID:      occurrenceID,
		Status:                 models.RegistrationStatusRegistered,
		StripePaymentIntentID:  "pi_test_123",
		StripeCustomerID:       "cus_test_123",
		OrgStripeAccountID:     "acct_test_123",
		StripePaymentMethodID:  "pm_test_123",
		TotalAmount:            10000,
		ProviderAmount:         8500,
		PlatformFeeAmount:      1500,
		Currency:               "usd",
		PaymentIntentStatus:    "requires_capture",
	}

	created, createErr := repo.CreateRegistration(ctx, createInput)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	updateInput := &models.UpdateRegistrationInput{
		ID: created.Body.ID,
	}
	updateInput.Body.ChildID = newChildID

	updated, updateErr := repo.UpdateRegistration(ctx, updateInput)
	require.Nil(t, updateErr)
	require.NotNil(t, updated)
	assert.Equal(t, newChildID, updated.Body.ChildID)
	assert.Equal(t, created.Body.GuardianID, updated.Body.GuardianID)
	assert.Equal(t, created.Body.EventOccurrenceID, updated.Body.EventOccurrenceID)
	assert.Equal(t, created.Body.Status, updated.Body.Status)
	assert.NotEmpty(t, updated.Body.EventName)
	assert.NotZero(t, updated.Body.OccurrenceStartTime)
	// Verify payment fields remain unchanged
	assert.Equal(t, created.Body.StripePaymentIntentID, updated.Body.StripePaymentIntentID)
	assert.Equal(t, created.Body.TotalAmount, updated.Body.TotalAmount)
	assert.Equal(t, created.Body.Currency, updated.Body.Currency)

	getInput := &models.GetRegistrationByIDInput{
		ID: created.Body.ID,
	}
	fetched, getErr := repo.GetRegistrationByID(ctx, getInput)
	require.Nil(t, getErr)
	assert.Equal(t, newChildID, fetched.Body.ChildID)
}

func TestUpdateRegistration_InvalidChildID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	childID := child.CreateTestChild(t, ctx, testDB).ID
	guardianID := guardian.CreateTestGuardian(t, ctx, testDB).ID
	occurrenceID := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB).ID

	createInput := &models.CreateRegistrationWithPaymentData{
		ChildID:                childID,
		GuardianID:             guardianID,
		EventOccurrenceID:      occurrenceID,
		Status:                 models.RegistrationStatusRegistered,
		StripePaymentIntentID:  "pi_test_123",
		StripeCustomerID:       "cus_test_123",
		OrgStripeAccountID:     "acct_test_123",
		StripePaymentMethodID:  "pm_test_123",
		TotalAmount:            10000,
		ProviderAmount:         8500,
		PlatformFeeAmount:      1500,
		Currency:               "usd",
		PaymentIntentStatus:    "requires_capture",
	}

	created, createErr := repo.CreateRegistration(ctx, createInput)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	invalidChildID := uuid.New()
	updateInput := &models.UpdateRegistrationInput{
		ID: created.Body.ID,
	}
	updateInput.Body.ChildID = invalidChildID

	updated, updateErr := repo.UpdateRegistration(ctx, updateInput)
	require.NotNil(t, updateErr)
	assert.Nil(t, updated)
}

func TestUpdateRegistration_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	nonExistentID := uuid.New()
	childID := child.CreateTestChild(t, ctx, testDB).ID
	
	updateInput := &models.UpdateRegistrationInput{
		ID: nonExistentID,
	}
	updateInput.Body.ChildID = childID

	updated, err := repo.UpdateRegistration(ctx, updateInput)

	require.NotNil(t, err)
	assert.Nil(t, updated)
}