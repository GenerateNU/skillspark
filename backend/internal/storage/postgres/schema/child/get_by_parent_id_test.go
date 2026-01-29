package child

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

	"skillspark/internal/storage/postgres/schema/guardian"
	"skillspark/internal/storage/postgres/schema/school"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChildRepository_GetChildrenByParentID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	p := guardian.CreateTestGuardian(t, ctx, testDB)
	s := school.CreateTestSchool(t, ctx, testDB)

	input1 := &models.CreateChildInput{}
	input1.Body.Name = "Child One"
	input1.Body.SchoolID = s.ID
	input1.Body.BirthMonth = 3
	input1.Body.BirthYear = 2018
	input1.Body.Interests = []string{"math"}
	input1.Body.GuardianID = p.ID

	child1, err := repo.CreateChild(ctx, input1)
	require.NoError(t, err)
	require.NotNil(t, child1)

	input2 := &models.CreateChildInput{}
	input2.Body.Name = "Child Two"
	input2.Body.SchoolID = s.ID
	input2.Body.BirthMonth = 7
	input2.Body.BirthYear = 2020
	input2.Body.Interests = []string{"science", "math"}
	input2.Body.GuardianID = p.ID

	child2, err := repo.CreateChild(ctx, input2)
	require.NoError(t, err)
	require.NotNil(t, child2)

	children, err := repo.GetChildrenByParentID(ctx, p.ID)

	require.NoError(t, err)
	require.NotNil(t, children)
	require.Len(t, children, 2)

	childIDs := map[uuid.UUID]bool{
		child1.ID: true,
		child2.ID: true,
	}

	for _, c := range children {
		require.True(t, childIDs[c.ID])
		require.Equal(t, p.ID, c.GuardianID)
	}
}

func TestChildRepository_GetChildrenByParentID_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	nonExistentID := uuid.New()

	child, err := repo.GetChildrenByParentID(ctx, nonExistentID)

	assert.Nil(t, child)
	assert.NotNil(t, err)
}
