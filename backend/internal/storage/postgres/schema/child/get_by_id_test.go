package child

import (
	"context"
	"testing"

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
