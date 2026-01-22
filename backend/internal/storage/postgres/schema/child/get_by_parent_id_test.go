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
	// this test should exist but cannot be written correctly until one can create a guardian on the fly to test for this.
	// otherwise we keep running into issues where more children were created under the same guardian and this test isn't really correct
}

func TestChildRepository_GetChildrenByParentID_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()

	nonExistentID := uuid.New()

	child, err := repo.GetChildrenByParentID(ctx, nonExistentID)

	assert.Nil(t, child)
	assert.NotNil(t, err)
}
