package child

import (
	"context"
	"net/http"
	"testing"

	"skillspark/internal/errs"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestChildRepository_DeleteChildByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()

	createdChild := CreateTestChild(t, ctx, repo)

	child, err := repo.DeleteChildByID(ctx, createdChild.ID)

	assert.Nil(t, err)
	assert.NotNil(t, child)

	child, err = repo.GetChildByID(ctx, createdChild.ID)

	assert.Nil(t, child)
	assert.NotNil(t, err)

}

func TestChildRepository_DeleteChildByID_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()

	nonExistentID := uuid.MustParse("00000000-0000-0000-0000-000000000000")

	child, err := repo.DeleteChildByID(ctx, nonExistentID)

	assert.Nil(t, child)
	assert.NotNil(t, err)
	httpErr, ok := err.(*errs.HTTPError)
	assert.True(t, ok, "expected *errs.HTTPError")

	assert.Equal(t, http.StatusNotFound, httpErr.Code)
}
