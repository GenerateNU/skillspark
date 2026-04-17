package review

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/child"
	"skillspark/internal/storage/postgres/schema/event"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"
	"skillspark/internal/storage/postgres/schema/registration"
	"skillspark/internal/storage/postgres/testutil"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAggregateReviews(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewReviewRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	regRepo := registration.NewRegistrationRepository(testDB)
	child := child.CreateTestChild(t, ctx, testDB)
	eo := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)

	input := &models.CreateRegistrationData{
		ChildID:           child.ID,
		GuardianID:        child.GuardianID,
		EventOccurrenceID: eo.ID,
		Status:            models.RegistrationStatusRegistered,
	}

	regOutput, _ := regRepo.CreateRegistration(ctx, input)

	descriptionTH := "รีวิวทดสอบ"

	ratingsToCreate := []int{4, 4, 2, 2, 2}
	for i, rating := range ratingsToCreate {
		reviewInput := &models.CreateReviewDBInput{}
		reviewInput.Body.RegistrationID = regOutput.Body.ID
		reviewInput.Body.GuardianID = &child.GuardianID
		reviewInput.Body.Rating = rating
		reviewInput.Body.Description_EN = "Review number " + strconv.Itoa(i+1)
		reviewInput.Body.Description_TH = &descriptionTH
		reviewInput.Body.Categories = []string{"interesting"}

		r, err := repo.CreateReview(ctx, reviewInput)
		require.Nil(t, err)
		require.NotNil(t, r)
	}

	aggregate, err := repo.GetAggregateReviews(ctx, eo.Event.ID)
	require.Nil(t, err)
	require.NotNil(t, aggregate)

	assert.Equal(t, eo.Event.ID, aggregate.EventID)
	assert.Equal(t, 5, aggregate.TotalReviews)
	assert.InDelta(t, 2.8, aggregate.AverageRating, 0.01)

	breakdown := make(map[int]int)
	for _, b := range aggregate.Breakdown {
		breakdown[int(b.Rating)] = b.ReviewCount
	}
	assert.Equal(t, 2, breakdown[4])
	assert.Equal(t, 3, breakdown[2])
}

func TestGetAggregateReviews_NoReviews(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewReviewRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	e := event.CreateTestEvent(t, ctx, testDB)

	aggregate, err := repo.GetAggregateReviews(ctx, e.ID)
	require.Nil(t, err)
	require.NotNil(t, aggregate)

	assert.Equal(t, e.ID, aggregate.EventID)
	assert.Equal(t, 0, aggregate.TotalReviews)
	assert.Equal(t, float64(0), aggregate.AverageRating)
	assert.Len(t, aggregate.Breakdown, 5)
	for _, b := range aggregate.Breakdown {
		assert.Equal(t, 0, b.ReviewCount)
	}
}
func TestGetAggregateReviews_AllSameRating(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewReviewRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	regRepo := registration.NewRegistrationRepository(testDB)
	child := child.CreateTestChild(t, ctx, testDB)
	eo := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)

	input := &models.CreateRegistrationData{
		ChildID:           child.ID,
		GuardianID:        child.GuardianID,
		EventOccurrenceID: eo.ID,
		Status:            models.RegistrationStatusRegistered,
	}

	regOutput, _ := regRepo.CreateRegistration(ctx, input)
	descriptionTH := "รีวิวทดสอบ"

	for i := 0; i < 4; i++ {
		reviewInput := &models.CreateReviewDBInput{}
		reviewInput.Body.RegistrationID = regOutput.Body.ID
		reviewInput.Body.GuardianID = &child.GuardianID
		reviewInput.Body.Rating = 5
		reviewInput.Body.Description_EN = "Perfect! " + strconv.Itoa(i+1)
		reviewInput.Body.Description_TH = &descriptionTH
		reviewInput.Body.Categories = []string{"interesting"}

		r, err := repo.CreateReview(ctx, reviewInput)
		require.Nil(t, err)
		require.NotNil(t, r)
	}

	aggregate, err := repo.GetAggregateReviews(ctx, eo.Event.ID)
	require.Nil(t, err)
	require.NotNil(t, aggregate)

	assert.Equal(t, 4, aggregate.TotalReviews)
	assert.InDelta(t, 5.0, aggregate.AverageRating, 0.01)
	assert.Len(t, aggregate.Breakdown, 5)
	assert.Equal(t, 0, aggregate.Breakdown[0].ReviewCount)
}
