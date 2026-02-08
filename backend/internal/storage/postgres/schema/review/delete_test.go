package review

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteReview_Success(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewReviewRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	review := CreateTestReview(t, ctx, testDB)

	err := repo.DeleteReview(ctx, review.ID)
	require.Nil(t, err)
}

func TestDeleteReview_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewReviewRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	fakeID := uuid.New()

	err := repo.DeleteReview(ctx, fakeID)
	require.NotNil(t, err)
	assert.Contains(t, err.Error(), "not found")
}
