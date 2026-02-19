package registration

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/child"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)
 
func TestCreateRegistration(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)

	child := child.CreateTestChild(t, ctx, testDB)
	occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)
	childID := child.ID
	guardianID := child.GuardianID
	occurrenceID := occurrence.ID

	require.NotNil(t, created)
	assert.NotEqual(t, uuid.Nil, created.ID)
	assert.NotEqual(t, uuid.Nil, created.ChildID)
	assert.NotEqual(t, uuid.Nil, created.GuardianID)
	assert.NotEqual(t, uuid.Nil, created.EventOccurrenceID)
	assert.Equal(t, models.RegistrationStatusRegistered, created.Status)
	assert.NotZero(t, created.CreatedAt)
	assert.NotZero(t, created.UpdatedAt)
	assert.NotEmpty(t, created.EventName)
	assert.NotZero(t, created.OccurrenceStartTime)
	
	// Verify payment fields
	assert.NotEmpty(t, created.StripePaymentIntentID)
	assert.NotEmpty(t, created.StripeCustomerID)
	assert.Equal(t, "acct_test_123", created.OrgStripeAccountID)
	assert.Equal(t, "pm_test_123", created.StripePaymentMethodID)
	assert.Equal(t, 10000, created.TotalAmount)
	assert.Equal(t, 8500, created.ProviderAmount)
	assert.Equal(t, 1500, created.PlatformFeeAmount)
	assert.Equal(t, "usd", created.Currency)
	assert.Equal(t, "requires_capture", created.PaymentIntentStatus)
	assert.Nil(t, created.CancelledAt)
	assert.Nil(t, created.PaidAt)
	assert.Equal(t, &childID, created.ChildID)
	assert.Equal(t, &guardianID, created.GuardianID)
	assert.Equal(t, occurrenceID, created.EventOccurrenceID)
	assert.Equal(t, models.RegistrationStatusRegistered, created.Status)
	assert.NotEqual(t, uuid.Nil, created.ID)
	assert.NotZero(t, created.CreatedAt)
	assert.NotZero(t, created.UpdatedAt)
	assert.NotEmpty(t, created.EventName)
	assert.NotZero(t, created.OccurrenceStartTime)
}

func TestCreateRegistration_VerifyEventNameJoin(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)

	require.NotNil(t, created)
	assert.NotEmpty(t, created.EventName)
}

func TestCreateRegistration_VerifyOccurrenceStartTime(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()
	t.Parallel()

	created := CreateTestRegistration(t, ctx, testDB)

	require.NotNil(t, created)
	assert.NotZero(t, created.OccurrenceStartTime)
}

func TestCreateRegistration_MultipleRegistrationsForSameChild(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	child := child.CreateTestChild(t, ctx, testDB)
	o1 := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)
	o2 := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)

	input1 := &models.CreateRegistrationWithPaymentData{
		ChildID:                child.ID,
		GuardianID:             child.GuardianID,
		EventOccurrenceID:      o1.ID,
		Status:                 models.RegistrationStatusRegistered,
		StripePaymentIntentID:  "pi_test_1",
		StripeCustomerID:       "cus_test_123",
		OrgStripeAccountID:     "acct_test_123",
		StripePaymentMethodID:  "pm_test_123",
		TotalAmount:            10000,
		ProviderAmount:         8500,
		PlatformFeeAmount:      1500,
		Currency:               "usd",
		PaymentIntentStatus:    "requires_capture",
	}

	created1, err := repo.CreateRegistration(ctx, input1)
	require.Nil(t, err)
	require.NotNil(t, created1)

	input2 := &models.CreateRegistrationWithPaymentData{
		ChildID:                child.ID,
		GuardianID:             child.GuardianID,
		EventOccurrenceID:      o2.ID,
		Status:                 models.RegistrationStatusRegistered,
		StripePaymentIntentID:  "pi_test_2",
		StripeCustomerID:       "cus_test_123",
		OrgStripeAccountID:     "acct_test_123",
		StripePaymentMethodID:  "pm_test_123",
		TotalAmount:            10000,
		ProviderAmount:         8500,
		PlatformFeeAmount:      1500,
		Currency:               "usd",
		PaymentIntentStatus:    "requires_capture",
	}

	created2, err := repo.CreateRegistration(ctx, input2)
	require.Nil(t, err)
	require.NotNil(t, created2)

	assert.NotEqual(t, created1.Body.ID, created2.Body.ID)
	assert.NotEqual(t, created1.Body.EventOccurrenceID, created2.Body.EventOccurrenceID)
}

func TestCreateRegistration_InvalidChildID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	child := child.CreateTestChild(t, ctx, testDB)
	occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)

	input := &models.CreateRegistrationWithPaymentData{
		ChildID:                uuid.New(),
		GuardianID:             child.GuardianID,
		EventOccurrenceID:      occurrence.ID,
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

	created, err := repo.CreateRegistration(ctx, input)

	require.NotNil(t, err)
	require.Nil(t, created)
}

func TestCreateRegistration_InvalidGuardianID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	child := child.CreateTestChild(t, ctx, testDB)
	occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)

	input := &models.CreateRegistrationWithPaymentData{
		ChildID:                child.ID,
		GuardianID:             uuid.New(),
		EventOccurrenceID:      occurrence.ID,
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

	created, err := repo.CreateRegistration(ctx, input)

	require.NotNil(t, err)
	require.Nil(t, created)
}

func TestCreateRegistration_InvalidEventOccurrenceID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	child := child.CreateTestChild(t, ctx, testDB)

	input := &models.CreateRegistrationWithPaymentData{
		ChildID:                child.ID,
		GuardianID:             child.GuardianID,
		EventOccurrenceID:      uuid.New(),
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

	created, err := repo.CreateRegistration(ctx, input)

	require.NotNil(t, err)
	require.Nil(t, created)
}

func TestCreateRegistration_IncreasesAttendeeCount(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	eventOccurrenceRepo := eventoccurrence.NewEventOccurrenceRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	child := child.CreateTestChild(t, ctx, testDB)
	occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)
	occurrenceID := occurrence.ID

	initialOccurrence, err := eventOccurrenceRepo.GetEventOccurrenceByID(ctx, occurrenceID)
	require.Nil(t, err)
	require.NotNil(t, initialOccurrence)
	initialCount := initialOccurrence.CurrEnrolled

	input := func() *models.CreateRegistrationWithPaymentData {
		i := &models.CreateRegistrationWithPaymentData{}
		i.ChildID = child.ID
		i.GuardianID = child.GuardianID
		i.EventOccurrenceID = occurrenceID
		i.Status = models.RegistrationStatusRegistered
		i.StripePaymentIntentID = "pi_test_123"
		i.StripeCustomerID = "cus_test_123"
		i.OrgStripeAccountID = "acct_test_123"
		i.StripePaymentMethodID = "pm_test_123"
		i.TotalAmount = 10000
		i.ProviderAmount = 8500
		i.PlatformFeeAmount = 1500
		i.Currency = "thb"
		i.PaymentIntentStatus = "requires_capture"
		return i
	}()

	created, err := repo.CreateRegistration(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, created)

	updatedOccurrence, err := eventOccurrenceRepo.GetEventOccurrenceByID(ctx, occurrenceID)
	require.Nil(t, err)
	require.NotNil(t, updatedOccurrence)
	assert.Equal(t, initialCount+1, updatedOccurrence.CurrEnrolled, "Attendee count should increase by 1 after successful registration")
}
