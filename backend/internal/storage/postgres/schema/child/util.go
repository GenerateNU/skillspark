package child

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/guardian"
	"skillspark/internal/storage/postgres/schema/school"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func CreateTestChild(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Child {

	t.Helper()
	repo := NewChildRepository(db)

	s := school.CreateTestSchool(t, ctx, db)
	g := guardian.CreateTestGuardian(t, ctx, db)

	input := &models.CreateChildInput{}
	input.Body.Name = "Test Child"
	input.Body.SchoolID = s.ID
	input.Body.BirthMonth = 5
	input.Body.BirthYear = 2019
	input.Body.Interests = []string{"math", "art"}
	input.Body.GuardianID = g.ID

	c, err := repo.CreateChild(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, c)

	return c
}
