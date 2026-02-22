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

	// get the total number of event occurrences in the test database
	var count int64
	row := testDB.QueryRow(ctx, "SELECT COUNT(*) FROM event_occurrence")
	sqlErr := row.Scan(&count)
	assert.Nil(t, sqlErr)

	// default pagination
	pagination := utils.NewPagination()

	// check that all 15 event occurrences in the test database are returned
	eventOccurrences, err := repo.GetAllEventOccurrences(ctx, pagination, models.GetAllEventOccurrencesFilter{})
	assert.Nil(t, err)
	assert.NotNil(t, eventOccurrences)
	assert.Equal(t, count, int64(len(eventOccurrences)))
}

func TestEventOccurrenceRepository_GetAllEventOccurrences_Pagination(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	// get the total number of event occurrences in the test database
	var count int64
	row := testDB.QueryRow(ctx, "SELECT COUNT(*) FROM event_occurrence")
	sqlErr := row.Scan(&count)
	assert.Nil(t, sqlErr)

	assert.GreaterOrEqual(t, count, int64(12))

	// test page 1 with limit 4
	pagination1 := utils.Pagination{Page: 1, Limit: 4}
	eventOccurrences1, err1 := repo.GetAllEventOccurrences(ctx, pagination1, models.GetAllEventOccurrencesFilter{})
	assert.Nil(t, err1)
	assert.NotNil(t, eventOccurrences1)
	assert.Equal(t, 4, len(eventOccurrences1))

	// test page 2 with limit 4
	pagination2 := utils.Pagination{Page: 2, Limit: 4}
	eventOccurrences2, err2 := repo.GetAllEventOccurrences(ctx, pagination2, models.GetAllEventOccurrencesFilter{})
	assert.Nil(t, err2)
	assert.NotNil(t, eventOccurrences2)
	assert.Equal(t, 4, len(eventOccurrences2))

	// test page 3 with limit 4
	pagination3 := utils.Pagination{Page: 3, Limit: 4}
	eventOccurrences3, err3 := repo.GetAllEventOccurrences(ctx, pagination3, models.GetAllEventOccurrencesFilter{})
	assert.Nil(t, err3)
	assert.NotNil(t, eventOccurrences3)
	assert.Equal(t, 4, len(eventOccurrences3))
}

func TestEventOccurrenceRepository_Filters_SearchDurationLocation(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	ctx := context.Background()
	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)

	pagination := utils.NewPagination()

	minDur := 60  // minutes
	maxDur := 120 // minutes

	eventOccurrences, err := repo.GetAllEventOccurrences(ctx, pagination, models.GetAllEventOccurrencesFilter{
		MinDurationMinutes: &minDur,
		MaxDurationMinutes: &maxDur,
	})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)

	for _, eo := range eventOccurrences {
		duration := int(eo.EndTime.Sub(eo.StartTime).Minutes())
		assert.GreaterOrEqual(t, duration, minDur)
		assert.LessOrEqual(t, duration, maxDur)
	}

	lat := 13.74
	lng := 100.545
	radiusKm := 5.0

	eventOccurrences, err = repo.GetAllEventOccurrences(ctx, pagination, models.GetAllEventOccurrencesFilter{
		Latitude:  &lat,
		Longitude: &lng,
		RadiusKm:  &radiusKm,
	})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)

	for _, eo := range eventOccurrences {
		dist := DistanceKm(eo.Location.Latitude, eo.Location.Longitude, lat, lng)
		assert.LessOrEqual(t, dist, radiusKm)
	}

	search := "Robotics"

	eventOccurrences, err = repo.GetAllEventOccurrences(ctx, pagination, models.GetAllEventOccurrencesFilter{
		Search: &search,
	})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)

	for _, eo := range eventOccurrences {
		title := eo.Event.Title
		desc := eo.Event.Description
		assert.True(t, containsIgnoreCase(title, search) || containsIgnoreCase(desc, search))
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

	eventOccurrences, err := repo.GetAllEventOccurrences(ctx, pagination, models.GetAllEventOccurrencesFilter{
		MinAge: &minAge,
		MaxAge: &maxAge,
	})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)

	for _, eo := range eventOccurrences {
		assert.True(t, *eo.Event.AgeRangeMin == 0 || *eo.Event.AgeRangeMin >= minAge)
		assert.True(t, *eo.Event.AgeRangeMax == 0 || *eo.Event.AgeRangeMax <= maxAge)
	}

	minDate := time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)
	maxDate := time.Date(2026, 2, 10, 23, 59, 59, 0, time.UTC)

	eventOccurrences, err = repo.GetAllEventOccurrences(ctx, pagination, models.GetAllEventOccurrencesFilter{
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

	eventOccurrences, err = repo.GetAllEventOccurrences(ctx, pagination, models.GetAllEventOccurrencesFilter{
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

	eventOccurrences, err = repo.GetAllEventOccurrences(ctx, pagination, models.GetAllEventOccurrencesFilter{
		SoldOut: &soldOut,
	})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)

	for _, eo := range eventOccurrences {
		assert.True(t, eo.CurrEnrolled >= eo.MaxAttendees)
	}

	soldOut = false

	eventOccurrences, err = repo.GetAllEventOccurrences(ctx, pagination, models.GetAllEventOccurrencesFilter{
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
