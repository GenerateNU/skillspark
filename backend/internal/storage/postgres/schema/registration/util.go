package registration

import (
	"context"
	"embed"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/child"
	"skillspark/internal/storage/postgres/schema/event"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

//go:embed sql/*.sql
var SqlRegistrationFiles embed.FS

func CreateTestRegistration(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Registration {
	t.Helper()

	repo := NewRegistrationRepository(db)

	child := child.CreateTestChild(t, ctx, db)
	occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, db)

	regInput := &models.CreateRegistrationData{
		AcceptLanguage:    "en-US",
		ChildID:           child.ID,
		GuardianID:        child.GuardianID,
		EventOccurrenceID: occurrence.ID,
		Status:            models.RegistrationStatusRegistered,
	}

	registration, err := repo.CreateRegistration(ctx, regInput)
	require.NoError(t, err)
	require.NotNil(t, registration.Body)

	paymentInput := &models.CreatePaymentData{
		RegistrationID:        registration.Body.ID,
		StripePaymentIntentID: "pi_test_" + child.ID.String()[:8],
		StripeCustomerID:      "cus_test_" + child.GuardianID.String()[:8],
		OrgStripeAccountID:    "acct_test_123",
		StripePaymentMethodID: "pm_test_123",
		TotalAmount:           10000, // $100.00
		ProviderAmount:        8500,  // $85.00
		PlatformFeeAmount:     1500,  // $15.00
		Currency:              "usd",
		PaymentIntentStatus:   "requires_capture",
	}

	err = repo.CreatePayment(ctx, paymentInput)
	require.NoError(t, err)

	// Re-fetch so the returned Registration has payment fields populated from the join
	full, err := repo.GetRegistrationByID(ctx, &models.GetRegistrationByIDInput{
		AcceptLanguage: "en-US",
		ID:             registration.Body.ID,
	}, nil)
	require.NoError(t, err)

	return &full.Body
}

// CreateTestRegistrationWithoutPayment creates a registered registration with no payment record.
// startTime controls the event occurrence start time, which must be in the future
// for GetRegistrationsForPaymentCreation to pick it up.
func CreateTestRegistrationWithoutPayment(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
	startTime time.Time,
) *models.Registration {
	t.Helper()

	repo := NewRegistrationRepository(db)
	eoRepo := eventoccurrence.NewEventOccurrenceRepository(db)

	c := child.CreateTestChild(t, ctx, db)
	e := event.CreateTestEvent(t, ctx, db)

	mid := uuid.MustParse("50000000-0000-0000-0000-000000000001")
	eoInput := &models.CreateEventOccurrenceInput{}
	eoInput.Body.ManagerId = &mid
	eoInput.Body.EventId = e.ID
	eoInput.Body.StartTime = startTime
	eoInput.Body.EndTime = startTime.Add(1 * time.Hour)
	eoInput.Body.MaxAttendees = 10
	eoInput.Body.Language = "en"

	occurrence, err := eoRepo.CreateEventOccurrence(ctx, eoInput)
	require.NoError(t, err)

	regInput := &models.CreateRegistrationData{
		AcceptLanguage:    "en-US",
		ChildID:           c.ID,
		GuardianID:        c.GuardianID,
		EventOccurrenceID: occurrence.ID,
		Status:            models.RegistrationStatusRegistered,
	}

	registration, err := repo.CreateRegistration(ctx, regInput)
	require.NoError(t, err)
	require.NotNil(t, registration.Body)

	full, err := repo.GetRegistrationByID(ctx, &models.GetRegistrationByIDInput{
		AcceptLanguage: "en-US",
		ID:             registration.Body.ID,
	}, nil)
	require.NoError(t, err)

	return &full.Body
}
