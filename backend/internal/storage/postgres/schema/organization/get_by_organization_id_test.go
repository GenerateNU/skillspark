package organization

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
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	// insert more event occurrences with the same organization id
	//

	// check that get by organization id returns multiple event occurrences with the same organization id
	eventOccurrences, err := repo.GetEventOccurrencesByOrganizationID(ctx, uuid.MustParse("40000000-0000-0000-0000-000000000001"))
	assert.Nil(t, err)
	assert.NotNil(t, eventOccurrences)
	for i := range eventOccurrences {
		assert.Equal(t, uuid.MustParse("40000000-0000-0000-0000-000000000001"), eventOccurrences[i].Event.OrganizationID)
	}
}