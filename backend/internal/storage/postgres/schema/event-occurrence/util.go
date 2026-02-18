package eventoccurrence

import (
	"context"
	"embed"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/event"
	"skillspark/internal/storage/postgres/schema/location"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

//go:embed sql/*.sql
var SqlEventOccurrenceFiles embed.FS

func CreateTestEventOccurrence(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.EventOccurrence {
	t.Helper()

	repo := NewEventOccurrenceRepository(db)

	e := event.CreateTestEvent(t, ctx, db)
	l := location.CreateTestLocation(t, ctx, db)

	mid := uuid.MustParse("50000000-0000-0000-0000-000000000001")
	start := time.Date(2026, time.February, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(2026, time.February, 1, 1, 0, 0, 0, time.Local)

	input := &models.CreateEventOccurrenceInput{}
	input.Body.ManagerId = &mid
	input.Body.EventId = e.ID
	input.Body.LocationId = l.ID
	input.Body.StartTime = start
	input.Body.EndTime = end
	input.Body.MaxAttendees = 10
	input.Body.Language = "en"

	eventOccurrence, err := repo.CreateEventOccurrence(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, eventOccurrence)

	return eventOccurrence
}
