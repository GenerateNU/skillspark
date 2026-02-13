package review

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/child"
	"skillspark/internal/storage/postgres/schema/event"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"
	"skillspark/internal/storage/postgres/schema/registration"
	"skillspark/internal/storage/postgres/testutil"
	"skillspark/internal/utils"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetReviewsByEventID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewReviewRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	regRepo := registration.NewRegistrationRepository(testDB)

	child := child.CreateTestChild(t, ctx, testDB)
	eo := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)

	input := func() *models.CreateRegistrationInput {
		i := &models.CreateRegistrationInput{}
		i.Body.ChildID = child.ID
		i.Body.GuardianID = child.GuardianID
		i.Body.EventOccurrenceID = eo.ID
		i.Body.Status = models.RegistrationStatusRegistered
		return i
	}()

	regoutput, _ := regRepo.CreateRegistration(ctx, input)

	var expectedReviews []*models.Review

	for i := 0; i < 3; i++ {

		input := &models.CreateReviewInput{}
		input.Body.RegistrationID = regoutput.Body.ID
		input.Body.GuardianID = child.GuardianID
		input.Body.Description = "Review number " + strconv.Itoa(i+2)
		input.Body.Categories = []string{"interesting"}

		r, err := repo.CreateReview(ctx, input)
		require.Nil(t, err)
		require.NotNil(t, r)

		expectedReviews = append(expectedReviews, r)
	}

	pagination := utils.Pagination{Limit: 10, Page: 1}
	reviews, err := repo.GetReviewsByEventID(ctx, eo.Event.ID, pagination)
	require.Nil(t, err)
	require.Len(t, reviews, len(expectedReviews))

	expectedIDs := make(map[uuid.UUID]bool)
	for _, r := range expectedReviews {
		expectedIDs[r.ID] = true
	}

	for _, r := range reviews {
		assert.Contains(t, expectedIDs, r.ID)
	}
}

func TestGetReviewsByEventID_NoReviews(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewReviewRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	e := event.CreateTestEvent(t, ctx, testDB)

	pagination := utils.Pagination{Limit: 10, Page: 1}
	reviews, err := repo.GetReviewsByEventID(ctx, e.ID, pagination)

	require.Nil(t, err)
	require.NotNil(t, reviews)
	assert.Len(t, reviews, 0)
}
