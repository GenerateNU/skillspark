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
	t.Parallel()

	midNew := uuid.MustParse("50000000-0000-0000-0000-000000000005")
	eid := uuid.MustParse("60000000-0000-0000-0000-00000000000e")
	lid := uuid.MustParse("10000000-0000-0000-0000-000000000008")
	startNew := time.Date(2027, time.March, 1, 9, 0, 0, 0, time.Local)
	endNew := time.Date(2027, time.March, 1, 11, 0, 0, 0, time.Local)
	max := 10
	lang := "th"
	curr := 8
	price := 75000
	currency := "thb"

	eventOccurrenceInput := func() *models.UpdateEventOccurrenceInput {
		input := &models.UpdateEventOccurrenceInput{}
		input.ID = uuid.MustParse("70000000-0000-0000-0000-000000000002")
		input.Body.ManagerId = &midNew
		input.Body.EventId = &eid
		input.Body.LocationId = &lid
		input.Body.StartTime = &startNew
		input.Body.EndTime = &endNew
		input.Body.MaxAttendees = &max
		input.Body.Language = &lang
		input.Body.CurrEnrolled = &curr
		input.Body.Price = &price
		input.Body.Currency = &currency
		return input
	}()

	eventOccurrence, err := repo.UpdateEventOccurrence(ctx, eventOccurrenceInput, nil)
	assert.NoError(t, err)
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
	assert.Equal(t, *eventOccurrenceInput.Body.Currency, eventOccurrence.Currency)

	// check updated event occurrence in database
	id := eventOccurrence.ID
	retrievedEventOccurrence, err := repo.GetEventOccurrenceByID(ctx, id, "en-US")
	assert.Nil(t, err)
	assert.NotNil(t, retrievedEventOccurrence)
	assert.Equal(t, &midNew, retrievedEventOccurrence.ManagerId)
	assert.Equal(t, eid, retrievedEventOccurrence.Event.ID)
	assert.Equal(t, lid, retrievedEventOccurrence.Location.ID)
	assert.Equal(t, startNew, retrievedEventOccurrence.StartTime)
	assert.Equal(t, endNew, retrievedEventOccurrence.EndTime)
	assert.Equal(t, 10, retrievedEventOccurrence.MaxAttendees)
	assert.Equal(t, "th", retrievedEventOccurrence.Language)
	assert.Equal(t, 8, retrievedEventOccurrence.CurrEnrolled)
	assert.Equal(t, 75000, retrievedEventOccurrence.Price)
	assert.Equal(t, "thb", retrievedEventOccurrence.Currency)
}