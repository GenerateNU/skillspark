package eventoccurrence

import (
	"context"
	"testing"
	"time"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEventOccurrenceRepository_UpdateEventOccurrence(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()

	mid_new := uuid.MustParse("50000000-0000-0000-0000-000000000005")
	eid := uuid.MustParse("60000000-0000-0000-0000-00000000000e")
	lid := uuid.MustParse("10000000-0000-0000-0000-000000000008")
	start_new := time.Date(0, time.December, 31, 19, 3, 58, 0, time.Local)
	end_new := time.Date(0, time.December, 31, 19, 3, 58, 0, time.Local)
	max := 10
	lang := "th"
	curr := 8
	price := 75000

	eventOccurrenceInput := func() *models.UpdateEventOccurrenceInput {
		input := &models.UpdateEventOccurrenceInput{}
		input.ID = uuid.MustParse("70000000-0000-0000-0000-000000000002")
		input.Body.ManagerId = &mid_new
		input.Body.EventId = &eid
		input.Body.LocationId = &lid
		input.Body.StartTime = &start_new
		input.Body.EndTime = &end_new
		input.Body.MaxAttendees = &max
		input.Body.Language = &lang
		input.Body.CurrEnrolled = &curr
		input.Body.Price = &price
		return input
	}()

	// check updated event occurrence struct
	eventOccurrence, err := repo.UpdateEventOccurrence(ctx, eventOccurrenceInput)
	assert.Nil(t, err)
	assert.NotNil(t, eventOccurrence)
	assert.Equal(t, eventOccurrenceInput.Body.ManagerId, eventOccurrence.ManagerId)
	assert.Equal(t, *eventOccurrenceInput.Body.EventId, eventOccurrence.Event.ID)
	assert.Equal(t, *eventOccurrenceInput.Body.LocationId, eventOccurrence.Location.ID)
	assert.Equal(t, *eventOccurrenceInput.Body.StartTime, eventOccurrence.StartTime)
	assert.Equal(t, *eventOccurrenceInput.Body.EndTime, eventOccurrence.EndTime)
	assert.Equal(t, *eventOccurrenceInput.Body.MaxAttendees, eventOccurrence.MaxAttendees)
	assert.Equal(t, *eventOccurrenceInput.Body.Language, eventOccurrence.Language)
	assert.Equal(t, *eventOccurrenceInput.Body.CurrEnrolled, eventOccurrence.CurrEnrolled)
	assert.Equal(t, *eventOccurrenceInput.Body.Price, eventOccurrence.Price)

	// check updated event occurrence in database
	id := eventOccurrence.ID
	retrievedEventOccurrence, err := repo.GetEventOccurrenceByID(ctx, id)
	assert.Nil(t, err)
	assert.NotNil(t, retrievedEventOccurrence)
	assert.Equal(t, &mid_new, retrievedEventOccurrence.ManagerId)
	assert.Equal(t, eid, retrievedEventOccurrence.Event.ID)
	assert.Equal(t, lid, retrievedEventOccurrence.Location.ID)
	assert.Equal(t, start_new, retrievedEventOccurrence.StartTime)
	assert.Equal(t, end_new, retrievedEventOccurrence.EndTime)
	assert.Equal(t, 10, retrievedEventOccurrence.MaxAttendees)
	assert.Equal(t, "th", retrievedEventOccurrence.Language)
	assert.Equal(t, 8, retrievedEventOccurrence.CurrEnrolled)
	assert.Equal(t, 75000, retrievedEventOccurrence.Price)
}