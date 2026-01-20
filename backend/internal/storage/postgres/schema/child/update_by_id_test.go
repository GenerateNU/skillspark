package child

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

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
