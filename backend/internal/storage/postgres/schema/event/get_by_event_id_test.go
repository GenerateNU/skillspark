package event

import (
	"context"
	"testing"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
)

func TestEventRepository_GetEventOccurrenceByEventId(t *testing.T){
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventRepository(testDB)
	ctx := context.Background()

	// check that get by event id returns multiple event occurrences with the same event id
	eventOccurrences, err := repo.GetEventOccurrencesByEventID(ctx, uuid.MustParse("60000000-0000-0000-0000-00000000000d"))
	assert.Nil(t, err)
	assert.NotNil(t, eventOccurrences)
	for i := range eventOccurrences {
		assert.Equal(t, uuid.MustParse("60000000-0000-0000-0000-00000000000d"), eventOccurrences[i].Event.ID)
	}
}