package saved

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/event"
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

	firstEvent := event.CreateTestEvent(t, ctx, testDB)
	secondEvent := event.CreateTestEvent(t, ctx, testDB)

	events := []*models.Event{
		firstEvent,
		secondEvent,
	}

	for i := 0; i < 2; i++ {

		input := &models.CreateSavedInput{}
		input.Body.EventID = events[i].ID
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
