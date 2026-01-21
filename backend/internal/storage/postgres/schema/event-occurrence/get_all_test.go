package eventoccurrence

import (
	"context"
	"testing"
	"skillspark/internal/storage/postgres/testutil"
	"skillspark/internal/utils"


	"github.com/stretchr/testify/assert"
)

func TestEventOccurrenceRepository_GetAllEventOccurrences(t *testing.T){
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()

	// default pagination
	pagination := utils.NewPagination()

	// check that all 15 event occurrences in the test database are returned
	eventOccurrences, err := repo.GetAllEventOccurrences(ctx, pagination)
	assert.Nil(t, err)
	assert.NotNil(t, eventOccurrences)
	assert.Equal(t, len(eventOccurrences), 15)
}

func TestEventOccurrenceRepository_GetAllEventOccurrences_Pagination(t *testing.T){
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventOccurrenceRepository(testDB)
	ctx := context.Background()

	// test page 1 with limit 4
	pagination1 := utils.Pagination{Page: 1, Limit: 4}
	eventOccurrences1, err1 := repo.GetAllEventOccurrences(ctx, pagination1)
	assert.Nil(t, err1)
	assert.NotNil(t, eventOccurrences1)
	assert.Equal(t, len(eventOccurrences1), 4)

	// test page 2 with limit 4
	pagination2 := utils.Pagination{Page: 2, Limit: 4}
	eventOccurrences2, err2 := repo.GetAllEventOccurrences(ctx, pagination2)
	assert.Nil(t, err2)
	assert.NotNil(t, eventOccurrences2)
	assert.Equal(t, len(eventOccurrences2), 4)

	// test page 3 with limit 4
	pagination3 := utils.Pagination{Page: 3, Limit: 4}
	eventOccurrences3, err3 := repo.GetAllEventOccurrences(ctx, pagination3)
	assert.Nil(t, err3)
	assert.NotNil(t, eventOccurrences3)
	assert.Equal(t, len(eventOccurrences3), 4)

	// test page 4 with limit 4
	pagination4 := utils.Pagination{Page: 4, Limit: 4}
	eventOccurrences4, err4 := repo.GetAllEventOccurrences(ctx, pagination4)
	assert.Nil(t, err4)
	assert.NotNil(t, eventOccurrences4)
	assert.Equal(t, len(eventOccurrences4), 1)
}