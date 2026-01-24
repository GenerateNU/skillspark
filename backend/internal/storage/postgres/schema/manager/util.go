package manager

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/organization"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
)

func CreateTestManager(
	t *testing.T,
	ctx context.Context,
) *models.Manager {
	t.Helper()

	testDB := testutil.SetupTestDB(t)
	repo := NewManagerRepository(testDB)

	org := organization.CreateTestOrganization(t, ctx)

	input := &models.CreateManagerInput{}
	input.Body.UserID = uuid.MustParse("f6a7b8c9-d0e1-4f2a-3b4c-5d6e7f8a9b0c")
	input.Body.OrganizationID = &org.ID
	input.Body.Role = "Assistant Manager"

	organization, _ := repo.CreateManager(ctx, input)

	return organization
}
