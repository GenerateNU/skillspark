package child

import (
	"context"
	"net/http"
	"testing"

	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"skillspark/internal/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestChildRepository_UpdateChildByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()

	testChild := CreateTestChild(t, ctx, repo)

	input := &models.UpdateChildInput{}
	name := "Updated Child"
	month := 11
	year := 2000
	interests := []string{"math"}

	// all pointers here
	input.Body.Name = &name
	input.Body.BirthMonth = &month
	input.Body.BirthYear = &year
	input.Body.Interests = &interests

	updatedChild, err := repo.UpdateChildByID(ctx, testChild.ID, input)

	assert.Nil(t, err)
	assert.NotNil(t, updatedChild)

	assert.Equal(t, "Updated Child", updatedChild.Name)
	assert.Equal(t, 11, updatedChild.BirthMonth)
	assert.Equal(t, 2000, updatedChild.BirthYear)
	assert.ElementsMatch(t, []string{"math"}, updatedChild.Interests)

	assert.Equal(t, testChild.GuardianID, updatedChild.GuardianID)
	assert.Equal(t, testChild.SchoolID, updatedChild.SchoolID)
	assert.Equal(t, testChild.SchoolName, updatedChild.SchoolName)

	assert.False(t, updatedChild.UpdatedAt.IsZero())
	assert.False(t, updatedChild.CreatedAt.IsZero())
}

func TestChildRepository_UpdateChildByID_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()

	nonExistentID := uuid.MustParse("00000000-0000-0000-0000-000000000000")

	input := &models.UpdateChildInput{}
	input.Body.Name = utils.PtrString("Updated Name")
	input.Body.BirthMonth = utils.PtrInt(5)
	input.Body.BirthYear = utils.PtrInt(2019)

	child, err := repo.UpdateChildByID(ctx, nonExistentID, input)

	assert.Nil(t, child)
	assert.NotNil(t, err)

	httpErr, ok := err.(*errs.HTTPError)
	assert.True(t, ok, "expected *errs.HTTPError")

	assert.Equal(t, http.StatusNotFound, httpErr.Code)
}
