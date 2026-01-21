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

func TestChildRepository_GetChildByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()

	testChild := CreateTestChild(t, ctx, repo)

	child, err := repo.GetChildByID(ctx, testChild.ID)

	assert.Nil(t, err)
	assert.NotNil(t, child)

	assert.NotEqual(t, uuid.Nil, child.ID)
	assert.Equal(t, testChild.SchoolID, child.SchoolID)
	assert.Equal(t, testChild.GuardianID, child.GuardianID)

	assert.Equal(t, testChild.Name, child.Name)
	assert.Equal(t, testChild.BirthMonth, child.BirthMonth)
	assert.Equal(t, testChild.BirthYear, child.BirthYear)

	assert.ElementsMatch(
		t,
		testChild.Interests,
		child.Interests,
	)

	assert.NotEmpty(t, child.SchoolName)

	assert.False(t, child.CreatedAt.IsZero())
	assert.False(t, child.UpdatedAt.IsZero())
}

func TestChildRepository_GetChildByID_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()

	nonExistentID := uuid.MustParse("00000000-0000-0000-0000-000000000000")

	child, err := repo.GetChildByID(ctx, nonExistentID)

	assert.Nil(t, child)
	assert.NotNil(t, err)
	httpErr, ok := err.(*errs.HTTPError)
	assert.True(t, ok, "expected *errs.HTTPError")

	assert.Equal(t, http.StatusNotFound, httpErr.Code)
}
