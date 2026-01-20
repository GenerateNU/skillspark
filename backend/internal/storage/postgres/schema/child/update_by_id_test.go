package child

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	interests := []models.Interest{"walking"}

	// all pointers here
	input.Body.Name = &name
	input.Body.BirthMonth = &month
	input.Body.BirthYear = &year
	input.Body.Interests = &interests

	updatedChild, err := repo.UpdateChildByID(ctx, testChild.ID, input)

	require.NoError(t, err)
	require.NotNil(t, updatedChild)

	assert.Equal(t, "Updated Child", updatedChild.Name)
	assert.Equal(t, 11, updatedChild.BirthMonth)
	assert.Equal(t, 2000, updatedChild.BirthYear)
	assert.ElementsMatch(t, []models.Interest{"walking"}, updatedChild.Interests)

	assert.Equal(t, testChild.GuardianID, updatedChild.GuardianID)
	assert.Equal(t, testChild.SchoolID, updatedChild.SchoolID)
	assert.Equal(t, testChild.SchoolName, updatedChild.SchoolName)

	assert.False(t, updatedChild.UpdatedAt.IsZero())
	assert.False(t, updatedChild.CreatedAt.IsZero())
}
