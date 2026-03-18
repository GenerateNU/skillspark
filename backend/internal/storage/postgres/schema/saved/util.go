package saved

import (
	"context"
	"embed"
	"skillspark/internal/models"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"
	"skillspark/internal/storage/postgres/schema/guardian"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

//go:embed sql/*.sql
var SqlSavedFiles embed.FS

func CreateTestSaved(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Saved {

	t.Helper()

	repo := NewSavedRepository(db)

	e := eventoccurrence.CreateTestEventOccurrence(t, ctx, db)
	g := guardian.CreateTestGuardian(t, ctx, db)

	input := &models.CreateSavedInput{}
	input.Body.EventOccurrenceID = e.ID
	input.Body.GuardianID = g.ID

	c, err := repo.CreateSaved(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, c)

	return c
}
