package child

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func CreateTestChildren(
	t *testing.T,
	ctx context.Context,
	repo *ChildRepository,
) []*models.Child {
	t.Helper()

	schoolID, err := uuid.Parse("20000000-0000-0000-0000-000000000001")
	assert.Nil(t, err)

	guardianID, err := uuid.Parse("88888888-8888-8888-8888-888888888888")
	assert.Nil(t, err)

	input1 := &models.CreateChildInput{}
	input1.Body.Name = "Test Child 1"
	input1.Body.SchoolID = schoolID
	input1.Body.BirthMonth = 5
	input1.Body.BirthYear = 2019
	input1.Body.Interests = []string{"math"}
	input1.Body.GuardianID = guardianID

	child1, err := repo.CreateChild(ctx, input1)
	assert.Nil(t, err)
	assert.NotNil(t, child1)

	input2 := &models.CreateChildInput{}
	input2.Body.Name = "Test Child 2"
	input2.Body.SchoolID = schoolID
	input2.Body.BirthMonth = 8
	input2.Body.BirthYear = 2021
	input2.Body.Interests = []string{"science"}
	input2.Body.GuardianID = guardianID

	child2, err := repo.CreateChild(ctx, input2)
	assert.Nil(t, err)
	assert.NotNil(t, child2)

	return []*models.Child{child1, child2}
}

func TestChildRepository_GetChildrenByParentID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()

	created := CreateTestChildren(t, ctx, repo)
	child1 := created[0]
	child2 := created[1]

	children, err := repo.GetChildrenByParentID(ctx, child1.GuardianID)
	assert.Nil(t, err)
	assert.NotNil(t, children)

	childMap := make(map[uuid.UUID]models.Child)
	for _, c := range children {
		childMap[c.ID] = c
	}

	c1, ok := childMap[child1.ID]
	assert.Equal(t, ok, true)
	assert.Equal(t, child1.Name, c1.Name)
	assert.Equal(t, child1.SchoolID, c1.SchoolID)
	assert.Equal(t, child1.GuardianID, c1.GuardianID)
	assert.ElementsMatch(t, child1.Interests, c1.Interests)
	assert.NotEmpty(t, c1.SchoolName)

	c2, ok := childMap[child2.ID]
	assert.Equal(t, ok, true)
	assert.Equal(t, child2.Name, c2.Name)
	assert.Equal(t, child2.SchoolID, c2.SchoolID)
	assert.Equal(t, child2.GuardianID, c2.GuardianID)
	assert.ElementsMatch(t, child2.Interests, c2.Interests)
	assert.NotEmpty(t, c2.SchoolName)
}
