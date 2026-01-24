package child

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func CreateTestChild(
	t *testing.T,
	ctx context.Context,
) *models.Child {
	t.Helper()

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)

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
