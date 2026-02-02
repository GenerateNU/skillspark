package organization

import (
	"context"
	"skillspark/internal/models"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func CreateTestOrganization(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Organization {
	t.Helper()

	repo := NewOrganizationRepository(db)

	active := true
	i := &models.CreateOrganizationInput{}
	i.Body.Name = "Test Corp"
	i.Body.Active = &active

	organization, err := repo.CreateOrganization(ctx, i)

	require.NoError(t, err)
	require.NotNil(t, organization)

	return organization
}
