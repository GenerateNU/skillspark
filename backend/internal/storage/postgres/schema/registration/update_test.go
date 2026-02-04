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

func TestUpdateRegistration(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	child := child.CreateTestChild(t, ctx, testDB)
	childID := child.ID
	guardianID := child.GuardianID
	occurrenceID := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB).ID

	createInput := func() *models.CreateRegistrationInput {
		i := &models.CreateRegistrationInput{}
		i.Body.ChildID = childID
		i.Body.GuardianID = guardianID
		i.Body.EventOccurrenceID = occurrenceID
		i.Body.Status = models.RegistrationStatusRegistered
		return i
	}()

	created, createErr := repo.CreateRegistration(ctx, createInput)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	newStatus := models.RegistrationStatusCancelled
	updateInput := &models.UpdateRegistrationInput{
		ID: created.Body.ID,
	}
	updateInput.Body.Status = &newStatus

	updated, updateErr := repo.UpdateRegistration(ctx, updateInput)
	require.Nil(t, updateErr)
	require.NotNil(t, updated)
	assert.Equal(t, models.RegistrationStatusCancelled, updated.Body.Status)
	assert.Equal(t, created.Body.ChildID, updated.Body.ChildID)
	assert.Equal(t, created.Body.GuardianID, updated.Body.GuardianID)
	assert.Equal(t, created.Body.EventOccurrenceID, updated.Body.EventOccurrenceID)
	assert.NotEmpty(t, updated.Body.EventName)
	assert.NotZero(t, updated.Body.OccurrenceStartTime)

	getInput := &models.GetRegistrationByIDInput{
		ID: created.Body.ID,
	}
	fetched, getErr := repo.GetRegistrationByID(ctx, getInput)
	require.Nil(t, getErr)
	assert.Equal(t, models.RegistrationStatusCancelled, fetched.Body.Status)
}

func TestUpdateRegistration_ChangeEventOccurrence(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	child := child.CreateTestChild(t, ctx, testDB)
	childID := child.ID
	guardianID := child.GuardianID
	occurrenceID1 := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB).ID

	createInput := func() *models.CreateRegistrationInput {
		i := &models.CreateRegistrationInput{}
		i.Body.ChildID = childID
		i.Body.GuardianID = guardianID
		i.Body.EventOccurrenceID = occurrenceID1
		i.Body.Status = models.RegistrationStatusRegistered
		return i
	}()

	created, createErr := repo.CreateRegistration(ctx, createInput)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	newOccurrenceID := uuid.MustParse("70000000-0000-0000-0000-000000000003")
	updateInput := &models.UpdateRegistrationInput{
		ID: created.Body.ID,
	}
	updateInput.Body.EventOccurrenceID = &newOccurrenceID

	updated, updateErr := repo.UpdateRegistration(ctx, updateInput)
	require.Nil(t, updateErr)
	require.NotNil(t, updated)
	assert.Equal(t, newOccurrenceID, updated.Body.EventOccurrenceID)
	assert.NotEmpty(t, updated.Body.EventName)
}

func TestUpdateRegistration_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	nonExistentID := uuid.New()
	newStatus := models.RegistrationStatusCancelled
	updateInput := &models.UpdateRegistrationInput{
		ID: nonExistentID,
	}
	updateInput.Body.Status = &newStatus

	updated, err := repo.UpdateRegistration(ctx, updateInput)

	require.NotNil(t, err)
	assert.Nil(t, updated)
}
