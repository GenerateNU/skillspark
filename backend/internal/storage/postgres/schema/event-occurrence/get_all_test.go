package eventoccurrence

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"skillspark/internal/utils"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEventOccurrenceRepository_GetAllEventOccurrences(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	var count int64
	row := testDB.QueryRow(ctx, "SELECT COUNT(*) FROM event_occurrence")
	err := row.Scan(&count)
	require.NoError(t, err)

	pagination := utils.NewPagination()

	// check that all 15 event occurrences in the test database are returned
	eventOccurrences, err := repo.GetAllEventOccurrences(ctx, pagination, "en-US", models.GetAllEventOccurrencesFilter{})
	assert.Nil(t, err)
	assert.NotNil(t, eventOccurrences)
	assert.Equal(t, count, int64(len(eventOccurrences)))

	for _, eo := range eventOccurrences {
		assert.Greater(t, eo.Price, 0)
		assert.Equal(t, "thb", eo.Currency)
	}
}

func TestEventOccurrenceRepository_GetAllEventOccurrences_Pagination(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	var count int64
	row := testDB.QueryRow(ctx, "SELECT COUNT(*) FROM event_occurrence")
	err := row.Scan(&count)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, count, int64(12))

	pagination1 := utils.Pagination{Page: 1, Limit: 4}
	eventOccurrences1, err1 := repo.GetAllEventOccurrences(ctx, pagination1, "en-US", models.GetAllEventOccurrencesFilter{})
	assert.Nil(t, err1)
	assert.NotNil(t, eventOccurrences1)
	assert.Equal(t, 4, len(eventOccurrences1))
	for _, eo := range eventOccurrences1 {
		assert.Greater(t, eo.Price, 0)
		assert.Equal(t, "thb", eo.Currency)
	}

	pagination2 := utils.Pagination{Page: 2, Limit: 4}
	eventOccurrences2, err2 := repo.GetAllEventOccurrences(ctx, pagination2, "en-US", models.GetAllEventOccurrencesFilter{})
	assert.Nil(t, err2)
	assert.NotNil(t, eventOccurrences2)
	assert.Equal(t, 4, len(eventOccurrences2))
	for _, eo := range eventOccurrences2 {
		assert.Greater(t, eo.Price, 0)
		assert.Equal(t, "thb", eo.Currency)
	}

	pagination3 := utils.Pagination{Page: 3, Limit: 4}
	eventOccurrences3, err3 := repo.GetAllEventOccurrences(ctx, pagination3, "en-US", models.GetAllEventOccurrencesFilter{})
	assert.Nil(t, err3)
	assert.NotNil(t, eventOccurrences3)
	assert.Equal(t, 4, len(eventOccurrences3))
	for _, eo := range eventOccurrences3 {
		assert.Greater(t, eo.Price, 0)
		assert.Equal(t, "thb", eo.Currency)
	}

	// verify pages don't overlap
	ids1 := make(map[string]bool)
	for _, eo := range eventOccurrences1 {
		ids1[eo.ID.String()] = true
	}
	for _, eo := range eventOccurrences2 {
		assert.False(t, ids1[eo.ID.String()], "page 2 should not contain items from page 1")
	}
}

func TestEventOccurrenceRepository_Filters_SearchDurationLocation(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	ctx := context.Background()
	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	t.Parallel()

	pagination := utils.NewPagination()

	minDur := 60
	maxDur := 120

	eventOccurrences, err := repo.GetAllEventOccurrences(ctx, pagination, "en-US", models.GetAllEventOccurrencesFilter{
		MinDurationMinutes: &minDur,
		MaxDurationMinutes: &maxDur,
	})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)
	for _, eo := range eventOccurrences {
		duration := int(eo.EndTime.Sub(eo.StartTime).Minutes())
		assert.GreaterOrEqual(t, duration, minDur)
		assert.LessOrEqual(t, duration, maxDur)
		assert.Equal(t, "thb", eo.Currency)
	}

	lat := 13.74
	lng := 100.545
	radiusKm := 5.0

	eventOccurrences, err = repo.GetAllEventOccurrences(ctx, pagination, "en-US", models.GetAllEventOccurrencesFilter{
		Latitude:  &lat,
		Longitude: &lng,
		RadiusKm:  &radiusKm,
	})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)
	for _, eo := range eventOccurrences {
		dist := DistanceKm(eo.Location.Latitude, eo.Location.Longitude, lat, lng)
		assert.LessOrEqual(t, dist, radiusKm)
		assert.Equal(t, "thb", eo.Currency)
	}

	search := "Robotics"

	eventOccurrences, err = repo.GetAllEventOccurrences(ctx, pagination, "en-US", models.GetAllEventOccurrencesFilter{
		Search: &search,
	})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)
	for _, eo := range eventOccurrences {
		assert.True(t, containsIgnoreCase(eo.Event.Title, search) || containsIgnoreCase(eo.Event.Description, search))
		assert.Equal(t, "thb", eo.Currency)
	}
}

func TestEventOccurrenceRepository_Filters_NewFilters(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	ctx := context.Background()
	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)

	pagination := utils.NewPagination()

	minAge := 10
	maxAge := 15

	eventOccurrences, err := repo.GetAllEventOccurrences(ctx, pagination, "en-US", models.GetAllEventOccurrencesFilter{
		MinAge: &minAge,
		MaxAge: &maxAge,
	})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)

	for _, eo := range eventOccurrences {
		// overlap: event's upper bound must reach filter's lower bound
		assert.True(t, eo.Event.AgeRangeMax == nil || *eo.Event.AgeRangeMax >= minAge)
		// overlap: event's lower bound must not exceed filter's upper bound
		assert.True(t, eo.Event.AgeRangeMin == nil || *eo.Event.AgeRangeMin <= maxAge)
	}

	minDate := time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)
	maxDate := time.Date(2026, 2, 10, 23, 59, 59, 0, time.UTC)

	eventOccurrences, err = repo.GetAllEventOccurrences(ctx, pagination, "en-US", models.GetAllEventOccurrencesFilter{
		MinDate: &minDate,
		MaxDate: &maxDate,
	})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)

	for _, eo := range eventOccurrences {
		assert.True(t, eo.StartTime.IsZero() || !eo.StartTime.Before(minDate))
		assert.True(t, eo.EndTime.IsZero() || !eo.EndTime.After(maxDate))
	}

	category := "science"

	eventOccurrences, err = repo.GetAllEventOccurrences(ctx, pagination, "en-US", models.GetAllEventOccurrencesFilter{
		Category: &category,
	})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)

	for _, eo := range eventOccurrences {
		found := false
		for _, c := range eo.Event.Category {
			if string(c) == category {
				found = true
				break
			}
		}
		assert.True(t, found, "event %s does not contain category %s", eo.Event.Title, category)
	}

	soldOut := true

	eventOccurrences, err = repo.GetAllEventOccurrences(ctx, pagination, "en-US", models.GetAllEventOccurrencesFilter{
		SoldOut: &soldOut,
	})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)

	for _, eo := range eventOccurrences {
		assert.True(t, eo.CurrEnrolled >= eo.MaxAttendees)
	}

	soldOut = false

	eventOccurrences, err = repo.GetAllEventOccurrences(ctx, pagination, "en-US", models.GetAllEventOccurrencesFilter{
		SoldOut: &soldOut,
	})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)

	for _, eo := range eventOccurrences {
		assert.True(t, eo.CurrEnrolled < eo.MaxAttendees)
	}
}

// helper for case-insensitive search match
func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
