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

	child1 := child.CreateTestChild(t, ctx, testDB)
	child2 := child.CreateTestChild(t, ctx, testDB)
	childID := child1.ID
	guardianID := child1.GuardianID
	occurrenceID := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB).ID

	createInput := &models.CreateRegistrationData{
		ChildID:           childID,
		GuardianID:        guardianID,
		EventOccurrenceID: occurrenceID,
		Status:            models.RegistrationStatusRegistered,
	}

	created, createErr := repo.CreateRegistration(ctx, createInput)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	updateInput := &models.UpdateRegistrationInput{
		AcceptLanguage: "en-US",
		ID:             created.Body.ID,
	}
	updateInput.Body.ChildID = &child2.ID

	updated, updateErr := repo.UpdateRegistration(ctx, updateInput)
	require.Nil(t, updateErr)
	require.NotNil(t, updated)
	assert.Equal(t, child2.ID, updated.Body.ChildID)
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
	fetched, getErr := repo.GetRegistrationByID(ctx, getInput, nil)
	require.Nil(t, getErr)
	assert.Equal(t, child2.ID, fetched.Body.ChildID)
}

func TestUpdateRegistration_InvalidChildID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	created := CreateTestRegistration(t, ctx, testDB)

	invalidChildID := uuid.New()
	updateInput := &models.UpdateRegistrationInput{
		ID: created.ID,
	}
	updateInput.Body.ChildID = &invalidChildID

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
	updateInput.Body.ChildID = &childID

	updated, err := repo.UpdateRegistration(ctx, updateInput)

	require.NotNil(t, err)
	assert.Nil(t, updated)
}

func TestUpdateRegistration_DecrementsAttendeeCount(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	eventOccurrenceRepo := eventoccurrence.NewEventOccurrenceRepository(testDB)
	ctx := context.Background()

	childID := child.CreateTestChild(t, ctx, testDB).ID
	guardianID := guardian.CreateTestGuardian(t, ctx, testDB).ID
	occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)
	occurrenceID := occurrence.ID

	createInput := &models.CreateRegistrationData{
		ChildID:           childID,
		GuardianID:        guardianID,
		EventOccurrenceID: occurrenceID,
		Status:            models.RegistrationStatusRegistered,
	}

	created, createErr := repo.CreateRegistration(ctx, createInput)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	occurrenceAfterCreate, err := eventOccurrenceRepo.GetEventOccurrenceByID(ctx, occurrenceID, "en-US")
	require.Nil(t, err)
	require.NotNil(t, occurrenceAfterCreate)
	countAfterCreate := occurrenceAfterCreate.CurrEnrolled

	newStatus := models.RegistrationStatusCancelled
	updateInput := &models.UpdateRegistrationInput{
		ID: created.Body.ID,
	}
	updateInput.Body.Status = &newStatus

	updated, updateErr := repo.UpdateRegistration(ctx, updateInput)
	require.Nil(t, updateErr)
	require.NotNil(t, updated)
	assert.Equal(t, models.RegistrationStatusCancelled, updated.Body.Status)

	occurrenceAfterCancel, err := eventOccurrenceRepo.GetEventOccurrenceByID(ctx, occurrenceID, "en-US")
	require.Nil(t, err)
	require.NotNil(t, occurrenceAfterCancel)
	assert.Equal(t, countAfterCreate-1, occurrenceAfterCancel.CurrEnrolled, "Attendee count should decrease by 1 after cancelling registration")
}
