package eventoccurrence

import (
	"context"
	"math"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/event"
	"skillspark/internal/storage/postgres/schema/location"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

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

func DistanceKm(lat1, lng1, lat2, lng2 float64) float64 {
	const earthRadiusKm = 6371.0

	dLat := degreesToRadians(lat2 - lat1)
	dLng := degreesToRadians(lng2 - lng1)

	lat1Rad := degreesToRadians(lat1)
	lat2Rad := degreesToRadians(lat2)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLng/2)*math.Sin(dLng/2)*math.Cos(lat1Rad)*math.Cos(lat2Rad)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusKm * c
}

func degreesToRadians(deg float64) float64 {
	return deg * (math.Pi / 180)
}
