package manager

import (
	"context"
	"embed"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/organization"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

//go:embed sql/*.sql
var SqlManagerFiles embed.FS

func CreateTestManager(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Manager {
	t.Helper()

	repo := NewManagerRepository(db)

	org := organization.CreateTestOrganization(t, ctx, db)

	input := &models.CreateManagerInput{}
	input.Body.OrganizationID = &org.ID
	input.Body.Role = "Assistant Manager"

	manager, err := repo.CreateManager(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, manager)

	return manager
}
