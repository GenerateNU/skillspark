package eventoccurrence

import (
	"context"
	"testing"
	"time"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/event"
	"skillspark/internal/storage/postgres/schema/location"
	"skillspark/internal/storage/postgres/schema/manager"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/stretchr/testify/assert"
)

func TestEventOccurrenceRepository_DeleteEventOccurrence(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	eo := CreateTestEventOccurrence(t, ctx, testDB)

	err := repo.DeleteEventOccurrence(ctx, eo.ID)

	assert.NoError(t, err)

	_, err = repo.GetEventOccurrenceByID(ctx, eo.ID)
	assert.Error(t, err)

}

func TestEventOccurrenceRepository_DeleteEventOccurrence_Within24HoursFails(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()

	start := time.Now().Add(-2 * time.Hour)
	end := time.Now().Add(1 * time.Hour)

	eventOccurrenceInput := func() *models.CreateEventOccurrenceInput {
		input := &models.CreateEventOccurrenceInput{}
		input.Body.ManagerId = &manager.CreateTestManager(t, ctx, testDB).ID
		input.Body.EventId = event.CreateTestEvent(t, ctx, testDB).ID
		input.Body.LocationId = location.CreateTestLocation(t, ctx, testDB).ID
		input.Body.StartTime = start
		input.Body.EndTime = end
		input.Body.MaxAttendees = 10
		input.Body.Language = "en"
		return input
	}()

	eventOccurrence, err := repo.CreateEventOccurrence(ctx, eventOccurrenceInput)
	assert.NoError(t, err)
	assert.NotNil(t, eventOccurrence)

	err = repo.DeleteEventOccurrence(ctx, eventOccurrence.ID)

	assert.Error(t, err)

	_, getErr := repo.GetEventOccurrenceByID(ctx, eventOccurrence.ID)
	assert.NoError(t, getErr)
}
