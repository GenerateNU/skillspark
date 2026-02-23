package eventoccurrence

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"skillspark/internal/utils"
	"strings"
	"testing"

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

	eventOccurrences, err := repo.GetAllEventOccurrences(ctx, pagination, models.GetAllEventOccurrencesFilter{})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)
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
	eventOccurrences1, err := repo.GetAllEventOccurrences(ctx, pagination1, models.GetAllEventOccurrencesFilter{})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences1)
	assert.Equal(t, 4, len(eventOccurrences1))
	for _, eo := range eventOccurrences1 {
		assert.Greater(t, eo.Price, 0)
		assert.Equal(t, "thb", eo.Currency)
	}

	pagination2 := utils.Pagination{Page: 2, Limit: 4}
	eventOccurrences2, err := repo.GetAllEventOccurrences(ctx, pagination2, models.GetAllEventOccurrencesFilter{})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences2)
	assert.Equal(t, 4, len(eventOccurrences2))
	for _, eo := range eventOccurrences2 {
		assert.Greater(t, eo.Price, 0)
		assert.Equal(t, "thb", eo.Currency)
	}

	pagination3 := utils.Pagination{Page: 3, Limit: 4}
	eventOccurrences3, err := repo.GetAllEventOccurrences(ctx, pagination3, models.GetAllEventOccurrencesFilter{})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences3)
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
		assert.Equal(t, "thb", eo.Currency)
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
		assert.Equal(t, "thb", eo.Currency)
	}

	search := "Robotics"

	eventOccurrences, err = repo.GetAllEventOccurrences(ctx, pagination, models.GetAllEventOccurrencesFilter{
		Search: &search,
	})
	require.NoError(t, err)
	require.NotNil(t, eventOccurrences)
	for _, eo := range eventOccurrences {
		assert.True(t, containsIgnoreCase(eo.Event.Title, search) || containsIgnoreCase(eo.Event.Description, search))
		assert.Equal(t, "thb", eo.Currency)
	}
}

func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}