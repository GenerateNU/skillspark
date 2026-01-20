package child

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestChildren(
	t *testing.T,
	ctx context.Context,
	repo *ChildRepository,
) []*models.Child {
	t.Helper()

	schoolID, err := uuid.Parse("20000000-0000-0000-0000-000000000001")
	require.NoError(t, err)

	guardianID, err := uuid.Parse("88888888-8888-8888-8888-888888888888")
	require.NoError(t, err)

	input1 := &models.CreateChildInput{}
	input1.Body.Name = "Test Child 1"
	input1.Body.SchoolID = schoolID
	input1.Body.BirthMonth = 5
	input1.Body.BirthYear = 2019
	input1.Body.Interests = []models.Interest{"soccer"}
	input1.Body.GuardianID = guardianID

	child1, err := repo.CreateChild(ctx, input1)
	require.NoError(t, err)
	require.NotNil(t, child1)

	input2 := &models.CreateChildInput{}
	input2.Body.Name = "Test Child 2"
	input2.Body.SchoolID = schoolID
	input2.Body.BirthMonth = 8
	input2.Body.BirthYear = 2021
	input2.Body.Interests = []models.Interest{"climbing"}
	input2.Body.GuardianID = guardianID

	child2, err := repo.CreateChild(ctx, input2)
	require.NoError(t, err)
	require.NotNil(t, child2)

	return []*models.Child{child1, child2}
}

func TestChildRepository_GetChildrenByParentID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()

	// this creates two children under the same parent who also go to the same school
	created := CreateTestChildren(t, ctx, repo)
	child1 := created[0]
	child2 := created[1]

	children, err := repo.GetChildrenByParentID(ctx, child1.GuardianID)

	assert.NoError(t, err)
	// a better test for this necessitates the creation of new guardians for testing, which is an endpoint that does not yet exist
	assert.Len(t, children, 2)

	ids := []uuid.UUID{
		children[0].ID,
		children[1].ID,
	}

	assert.Contains(t, ids, child1.ID)
	assert.Contains(t, ids, child2.ID)

	assert.Equal(t, child1.GuardianID, children[0].GuardianID)
	assert.Equal(t, child1.GuardianID, children[1].GuardianID)

	assert.NotEmpty(t, children[0].SchoolName)
	assert.NotEmpty(t, children[1].SchoolName)
}
