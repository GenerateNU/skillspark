package organization

import (
	"context"
	"embed"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/location"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

//go:embed sql/*.sql
var SqlOrganizationFiles embed.FS

func CreateTestOrganization(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Organization {
	t.Helper()

	repo := NewOrganizationRepository(db)
	location := location.CreateTestLocation(t, ctx, db)

	active := true
	i := &models.CreateOrganizationInput{}
	i.Body.Name = "Test Corp"
	i.Body.Active = &active
	i.Body.LocationID = location.ID

	organization, err := repo.CreateOrganization(ctx, i, nil)
	require.NoError(t, err)
	require.NotNil(t, organization)

	return organization
}
