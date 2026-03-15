package saved

import (
	"context"
	"skillspark/internal/models"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"
	"skillspark/internal/storage/postgres/schema/guardian"
	"skillspark/internal/storage/postgres/testutil"
	"skillspark/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetReviewsByGuardianID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewSavedRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	firstSaved := CreateTestSaved(t, ctx, testDB)

	var expectedSaved []*models.Saved
	expectedSaved = append(expectedSaved, firstSaved)

	firstEO := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)
	secondEO := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)

	eventOccurrences := []*models.EventOccurrence{
		firstEO,
		secondEO,
	}

	for i := 0; i < 2; i++ {

		input := &models.CreateSavedInput{}
		input.Body.EventOccurrenceID = eventOccurrences[i].ID
		input.Body.GuardianID = firstSaved.GuardianID

		r, err := repo.CreateSaved(ctx, input)
		require.Nil(t, err)
		require.NotNil(t, r)

		expectedSaved = append(expectedSaved, r)
	}

	pagination := utils.Pagination{Limit: 10, Page: 1}
	reviews, err := repo.GetByGuardianID(ctx, firstSaved.GuardianID, pagination)
	require.Nil(t, err)
	require.Len(t, reviews, len(expectedSaved))
}

func TestGetReviewsByGuardianID_NoReviews(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewSavedRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	g := guardian.CreateTestGuardian(t, ctx, testDB)

	pagination := utils.Pagination{Limit: 10, Page: 1}
	reviews, err := repo.GetByGuardianID(ctx, g.ID, pagination)

	require.Nil(t, err)
	require.NotNil(t, reviews)
	assert.Len(t, reviews, 0)
}
