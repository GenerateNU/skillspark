package eventoccurrence

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/event"
	"skillspark/internal/storage/postgres/schema/location"
	"skillspark/internal/storage/postgres/schema/manager"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
	"time"
)

func CreateTestEventOccurrence(
	t *testing.T,
	ctx context.Context,
) *models.EventOccurrence {
	t.Helper()

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)

	e := event.CreateTestEvent(t, ctx)
	l := location.CreateTestLocation(t, ctx)
	m := manager.CreateTestManager(t, ctx)

	start := time.Date(2026, time.February, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(2026, time.February, 1, 1, 0, 0, 0, time.Local)

	input := &models.CreateEventOccurrenceInput{}
	input.Body.ManagerId = &m.ID
	input.Body.EventId = e.ID
	input.Body.LocationId = l.ID
	input.Body.StartTime = start
	input.Body.EndTime = end
	input.Body.MaxAttendees = 10
	input.Body.Language = "en"

	eventOccurrence, _ := repo.CreateEventOccurrence(ctx, input)

	return eventOccurrence
}
