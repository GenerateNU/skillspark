package saved

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
	repo := NewSavedRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	review := CreateTestSaved(t, ctx, testDB)

	err := repo.DeleteSaved(ctx, review.ID)
	require.Nil(t, err)
}

func TestDeleteSaved_NotFound(t *testing.T) {

	testDB := testutil.SetupTestDB(t)
	repo := NewSavedRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	fakeID := uuid.New()

	err := repo.DeleteSaved(ctx, fakeID)
	require.NotNil(t, err)
	assert.Contains(t, err.Error(), "not found")

}
