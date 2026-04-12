package eventoccurrence

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func bangkokInput() *models.GetTrendingEventOccurrencesInput {
	i := &models.GetTrendingEventOccurrencesInput{AcceptLanguage: "en-US"}
	i.Latitude = 13.74
	i.Longitude = 100.545
	i.Radius = 5
	i.MaxReturns = 5
	return i
}

func TestEventOccurrenceRepository_GetTrendingEventOccurrences(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	eventOccurrences, err := repo.GetTrendingEventOccurrences(ctx, bangkokInput())
	require.NoError(t, err)
	assert.NotNil(t, eventOccurrences)

	for _, eo := range eventOccurrences {
		assert.Greater(t, eo.Price, 0)
		assert.Equal(t, "thb", eo.Currency)
	}
}

// MaxReturns controls the LIMIT — verify it is respected
func TestEventOccurrenceRepository_GetTrendingEventOccurrences_LimitFive(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	eventOccurrences, err := repo.GetTrendingEventOccurrences(ctx, bangkokInput())
	require.NoError(t, err)
	assert.LessOrEqual(t, len(eventOccurrences), bangkokInput().MaxReturns,
		"trending should return at most MaxReturns results")
}

// MaxReturns=1 should return at most 1 result
func TestEventOccurrenceRepository_GetTrendingEventOccurrences_MaxReturnsOne(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	i := bangkokInput()
	i.MaxReturns = 1

	eventOccurrences, err := repo.GetTrendingEventOccurrences(ctx, i)
	require.NoError(t, err)
	assert.LessOrEqual(t, len(eventOccurrences), 1, "MaxReturns=1 should return at most 1 result")
}

// The SQL uses ROW_NUMBER() PARTITION BY event_id — only the next upcoming
// occurrence per event should appear, not multiple occurrences of the same event
func TestEventOccurrenceRepository_GetTrendingEventOccurrences_OneOccurrencePerEvent(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	eventOccurrences, err := repo.GetTrendingEventOccurrences(ctx, bangkokInput())
	require.NoError(t, err)

	seenEventIDs := make(map[string]bool)
	for _, eo := range eventOccurrences {
		eventID := eo.Event.ID.String()
		assert.False(t, seenEventIDs[eventID], "event %s appears more than once in trending results", eo.Event.Title)
		seenEventIDs[eventID] = true
	}
}

// The SQL filters ranked_occurrences to start_time > NOW() + INTERVAL '1 day'
// No past or same-day occurrences should appear
func TestEventOccurrenceRepository_GetTrendingEventOccurrences_OnlyFutureOccurrences(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	eventOccurrences, err := repo.GetTrendingEventOccurrences(ctx, bangkokInput())
	require.NoError(t, err)

	threshold := time.Now().Add(24 * time.Hour)
	for _, eo := range eventOccurrences {
		assert.True(t, eo.StartTime.After(threshold),
			"occurrence %s starts at %v which is within 1 day of now", eo.ID, eo.StartTime)
	}
}

// A large radius should return more results than a very small one
func TestEventOccurrenceRepository_GetTrendingEventOccurrences_RadiusAffectsResults(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	largeRadius := bangkokInput()
	largeRadius.Radius = 50
	largeRadius.MaxReturns = 100

	tinyRadius := bangkokInput()
	tinyRadius.Radius = 1
	tinyRadius.MaxReturns = 100

	large, err := repo.GetTrendingEventOccurrences(ctx, largeRadius)
	require.NoError(t, err)

	tiny, err := repo.GetTrendingEventOccurrences(ctx, tinyRadius)
	require.NoError(t, err)

	assert.GreaterOrEqual(t, len(large), len(tiny),
		"larger radius should return at least as many results as a smaller radius")
}

// A far-away coordinate should return no results since no orgs are nearby
func TestEventOccurrenceRepository_GetTrendingEventOccurrences_NoResultsForRemoteLocation(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	i := &models.GetTrendingEventOccurrencesInput{AcceptLanguage: "en-US"}
	i.Latitude = 51.5074 // London — no seed orgs nearby
	i.Longitude = -0.1278
	i.Radius = 5
	i.MaxReturns = 5

	eventOccurrences, err := repo.GetTrendingEventOccurrences(ctx, i)
	require.NoError(t, err)
	assert.Empty(t, eventOccurrences, "no trending results expected for a location with no nearby orgs")
}

// The SQL orders by total_enrolled DESC then start_time ASC — verify start_time
// is ascending when results share the same popularity (no recent enrollment history)
func TestEventOccurrenceRepository_GetTrendingEventOccurrences_OrderedByStartTime(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	eventOccurrences, err := repo.GetTrendingEventOccurrences(ctx, bangkokInput())
	require.NoError(t, err)

	if len(eventOccurrences) < 2 {
		t.Skip("Not enough results to verify ordering")
	}

	for i := 1; i < len(eventOccurrences); i++ {
		prev := eventOccurrences[i-1].StartTime
		curr := eventOccurrences[i].StartTime
		assert.False(t, curr.Before(prev),
			"occurrence at index %d (start: %v) should not come before index %d (start: %v)",
			i, curr, i-1, prev)
	}
}

// Verify all scanned fields are populated correctly from the joined tables
func TestEventOccurrenceRepository_GetTrendingEventOccurrences_ValidFields(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	eventOccurrences, err := repo.GetTrendingEventOccurrences(ctx, bangkokInput())
	require.NoError(t, err)
	require.NotEmpty(t, eventOccurrences)

	for _, eo := range eventOccurrences {
		// occurrence fields
		assert.NotEmpty(t, eo.ID)
		assert.False(t, eo.StartTime.IsZero())
		assert.False(t, eo.EndTime.IsZero())
		assert.True(t, eo.EndTime.After(eo.StartTime))
		assert.Greater(t, eo.MaxAttendees, 0)
		assert.GreaterOrEqual(t, eo.CurrEnrolled, 0)

		// event fields from JOIN
		assert.NotEmpty(t, eo.Event.ID)
		assert.NotEmpty(t, eo.Event.Title)
		assert.NotEmpty(t, eo.Event.OrganizationID)

		// location fields from JOIN on org
		assert.NotZero(t, eo.Location.Latitude)
		assert.NotZero(t, eo.Location.Longitude)
		assert.NotEmpty(t, eo.Location.Country)
	}
}
