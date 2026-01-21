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

func TestEventOccurrenceRepository_CreateEventOccurrence(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()

	eventOccurrenceInput := func() *models.CreateEventOccurrenceInput {
		input := &models.CreateEventOccurrenceInput{}
		input.Body.ManagerId = uuid.MustParse("50000000-0000-0000-0000-000000000001")
		input.Body.EventId = uuid.MustParse("60000000-0000-0000-0000-000000000001")
		input.Body.LocationId = uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
		start, _ := time.Parse(time.RFC3339, "2026-02-01T00:00:00Z")
		input.Body.StartTime = start
		end, _ := time.Parse(time.RFC3339, "2026-02-01T01:00:00Z")
		input.Body.EndTime = end
		input.Body.MaxAttendees = 10
		input.Body.Language = "en"
		return input
	}()

	// check created event occurrence struct
	eventOccurrence, err := repo.CreateEventOccurrence(ctx, eventOccurrenceInput)
	assert.Nil(t, err)
	assert.NotNil(t, eventOccurrence)
	assert.Equal(t, uuid.MustParse("50000000-0000-0000-0000-000000000001"), eventOccurrence.ManagerId)
	assert.Equal(t, uuid.MustParse("60000000-0000-0000-0000-000000000001"), eventOccurrence.Event.ID)
	assert.Equal(t, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"), eventOccurrence.Location.ID)
	assert.Equal(t, eventOccurrenceInput.Body.StartTime, eventOccurrence.StartTime)
	assert.Equal(t, eventOccurrenceInput.Body.EndTime, eventOccurrence.EndTime)
	assert.Equal(t, 10, eventOccurrence.MaxAttendees)
	assert.Equal(t, "en", eventOccurrence.Language)
	assert.Equal(t, 0, eventOccurrence.CurrEnrolled)

	// check created event occurrence in database
	id := eventOccurrence.ID
	retrievedEventOccurrence, err := repo.GetEventOccurrenceByID(ctx, id)
	assert.Nil(t, err)
	assert.NotNil(t, retrievedEventOccurrence)
	assert.Equal(t, uuid.MustParse("50000000-0000-0000-0000-000000000001"), retrievedEventOccurrence.ManagerId)
	assert.Equal(t, uuid.MustParse("60000000-0000-0000-0000-000000000001"), retrievedEventOccurrence.Event.ID)
	assert.Equal(t, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"), retrievedEventOccurrence.Location.ID)
	assert.Equal(t, eventOccurrenceInput.Body.StartTime, retrievedEventOccurrence.StartTime)
	assert.Equal(t, eventOccurrenceInput.Body.EndTime, retrievedEventOccurrence.EndTime)
	assert.Equal(t, 10, retrievedEventOccurrence.MaxAttendees)
	assert.Equal(t, "en", retrievedEventOccurrence.Language)
	assert.Equal(t, 0, retrievedEventOccurrence.CurrEnrolled)
}
