package child

import (
	"context"
	"skillspark/internal/models"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestChild(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Child {

	t.Helper()
	repo := NewChildRepository(db)

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

	c, err := repo.CreateChild(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, c)

	return c
}
