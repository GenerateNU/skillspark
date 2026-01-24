package organization

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
)

func CreateTestOrganization(
	t *testing.T,
	ctx context.Context,
) *models.Organization {
	t.Helper()

	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)

	active := true
	i := &models.CreateOrganizationInput{}
	i.Body.Name = "Test Corp"
	i.Body.Active = &active

	organization, _ := repo.CreateOrganization(ctx, i)

	return organization
}
