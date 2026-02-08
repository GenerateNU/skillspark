package review

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"skillspark/internal/utils"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetReviewsByGuardianID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewReviewRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	firstReview := CreateTestReview(t, ctx, testDB)

	var expectedReviews []*models.Review
	expectedReviews = append(expectedReviews, firstReview)

	for i := 0; i < 3; i++ {

		input := &models.CreateReviewInput{}
		input.Body.RegistrationID = firstReview.RegistrationID
		input.Body.GuardianID = firstReview.GuardianID
		input.Body.Description = "Review number " + strconv.Itoa(i+2)
		input.Body.Categories = []string{"interesting"}

		r, err := repo.CreateReview(ctx, input)
		require.Nil(t, err)
		require.NotNil(t, r)

		expectedReviews = append(expectedReviews, r)
	}

	pagination := utils.Pagination{Limit: 10, Page: 1}
	reviews, err := repo.GetReviewsByGuardianID(ctx, firstReview.GuardianID, pagination)
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
