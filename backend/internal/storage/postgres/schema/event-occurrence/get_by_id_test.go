package eventoccurrence

import (
	"context"
	"testing"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
)

func TestEventOccurrenceRepository_GetEventOccurrenceById(t *testing.T){
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()

	// check that get by id works for 3 different event occurrences
	eventOccurrence1, err := repo.GetEventOccurrenceByID(ctx, uuid.MustParse("70000000-0000-0000-0000-000000000001"))
	mid := uuid.MustParse("50000000-0000-0000-0000-000000000001")
	assert.Nil(t, err)
	assert.NotNil(t, eventOccurrence1)
	assert.Equal(t, &mid, eventOccurrence1.ManagerId)
	assert.Equal(t, uuid.MustParse("60000000-0000-0000-0000-000000000001"), eventOccurrence1.Event.ID)
	assert.Equal(t, uuid.MustParse("10000000-0000-0000-0000-000000000004"), eventOccurrence1.Location.ID)
	assert.Equal(t, 15, eventOccurrence1.MaxAttendees)
	assert.Equal(t, "en", eventOccurrence1.Language)
	assert.Equal(t, 8, eventOccurrence1.CurrEnrolled)

	eventOccurrence2, err := repo.GetEventOccurrenceByID(ctx, uuid.MustParse("70000000-0000-0000-0000-000000000003"))
	assert.Nil(t, err)
	assert.NotNil(t, eventOccurrence2)
	assert.Equal(t, &mid, eventOccurrence2.ManagerId)
	assert.Equal(t, uuid.MustParse("60000000-0000-0000-0000-000000000002"), eventOccurrence2.Event.ID)
	assert.Equal(t, uuid.MustParse("10000000-0000-0000-0000-000000000004"), eventOccurrence2.Location.ID)
	assert.Equal(t, 12, eventOccurrence2.MaxAttendees)
	assert.Equal(t, "en", eventOccurrence2.Language)
	assert.Equal(t, 10, eventOccurrence2.CurrEnrolled)

	eventOccurrence3, err := repo.GetEventOccurrenceByID(ctx, uuid.MustParse("70000000-0000-0000-0000-000000000002"))
	assert.Nil(t, err)
	assert.NotNil(t, eventOccurrence1)
	assert.Equal(t, &mid, eventOccurrence3.ManagerId)
	assert.Equal(t, uuid.MustParse("60000000-0000-0000-0000-000000000001"), eventOccurrence3.Event.ID)
	assert.Equal(t, uuid.MustParse("10000000-0000-0000-0000-000000000004"), eventOccurrence3.Location.ID)
	assert.Equal(t, 15, eventOccurrence3.MaxAttendees)
	assert.Equal(t, "en", eventOccurrence3.Language)
	assert.Equal(t, 5, eventOccurrence3.CurrEnrolled)

}