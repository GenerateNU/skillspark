package child

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// a test run of this way of testing, creating something
// before testing on it to ensure that it must exist
func CreateTestChild(
	t *testing.T,
	ctx context.Context,
	repo *ChildRepository,
) *models.Child {
	t.Helper()

	schoolID, err := uuid.Parse("20000000-0000-0000-0000-000000000001")
	assert.Nil(t, err)

	guardianID, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	assert.Nil(t, err)

	input := &models.CreateChildInput{}
	input.Body.Name = "Test Child"
	input.Body.SchoolID = schoolID
	input.Body.BirthMonth = 5
	input.Body.BirthYear = 2019
	input.Body.Interests = []string{"math", "art"}
	input.Body.GuardianID = guardianID

	child, _ := repo.CreateChild(ctx, input)
	return child
}

func TestChildRepository_CreateChild(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()

	sampleChild := CreateTestChild(t, ctx, repo)

	childInput := func() *models.CreateChildInput {
		input := &models.CreateChildInput{}
		input.Body.Name = "Alice"
		input.Body.SchoolID = sampleChild.SchoolID
		input.Body.BirthMonth = 5
		input.Body.BirthYear = 2019
		input.Body.Interests = []string{"math", "science"}
		input.Body.GuardianID = sampleChild.GuardianID
		return input
	}()

	child, err := repo.CreateChild(ctx, childInput)

	assert.Nil(t, err)
	assert.NotNil(t, child)

	assert.NotEqual(t, uuid.Nil, child.ID)
	assert.Equal(t, childInput.Body.SchoolID, child.SchoolID)
	assert.Equal(t, childInput.Body.GuardianID, child.GuardianID)

	assert.Equal(t, "Alice", child.Name)
	assert.Equal(t, 5, child.BirthMonth)
	assert.Equal(t, 2019, child.BirthYear)

	assert.ElementsMatch(
		t,
		[]string{"math", "science"},
		child.Interests,
	)

	assert.NotEmpty(t, child.SchoolName)

	assert.False(t, child.CreatedAt.IsZero())
	assert.False(t, child.UpdatedAt.IsZero())
}

func TestChildRepository_CreateChild_InvalidSchoolID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()

	guardianID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	input := &models.CreateChildInput{}
	input.Body.Name = "Invalid School Child"
	input.Body.SchoolID = uuid.New() // non-existent school
	input.Body.BirthMonth = 5
	input.Body.BirthYear = 2019
	input.Body.Interests = []string{"math"}
	input.Body.GuardianID = guardianID

	child, err := repo.CreateChild(ctx, input)

	assert.Nil(t, child)
	assert.Error(t, err)
}
