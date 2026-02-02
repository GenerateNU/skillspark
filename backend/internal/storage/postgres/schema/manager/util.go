package manager

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/organization"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func CreateTestManager(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Manager {
	t.Helper()

	repo := NewManagerRepository(db)

	org := organization.CreateTestOrganization(t, ctx, db)

	input := &models.CreateManagerInput{}
	input.Body.UserID = uuid.MustParse("f6a7b8c9-d0e1-4f2a-3b4c-5d6e7f8a9b0c")
	input.Body.OrganizationID = &org.ID
	input.Body.Role = "Assistant Manager"

	manager, err := repo.CreateManager(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, manager)

	return manager
}
