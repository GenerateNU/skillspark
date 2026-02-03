package organization

import (
	"context"
	"skillspark/internal/models"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
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

	organization, _ := repo.CreateOrganization(ctx, i)

	// t.Logf("error returned: %v", err)

	// require.NoError(t, err)
	// require.NotNil(t, organization)

	return organization
}
