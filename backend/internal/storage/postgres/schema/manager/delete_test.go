package manager

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/organization"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/stretchr/testify/assert"
)

func TestManagerRepository_Delete_Director(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewManagerRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	organizationID := organization.CreateTestOrganization(t, ctx, testDB).ID

	managerInput := func() *models.CreateManagerInput {
		input := &models.CreateManagerInput{}

		input.Body.OrganizationID = &organizationID
		input.Body.Name = "Delete Man"
		input.Body.Email = "delete.m@example.com"
		input.Body.Username = "delman"
		input.Body.LanguagePreference = "en"
		input.Body.Role = "Assistant Manager"
		return input
	}()
	createdManager, _ := repo.CreateManager(ctx, managerInput)
	manager, err := repo.DeleteManager(ctx, createdManager.ID)

	assert.Nil(t, err)
	assert.NotNil(t, manager)
	assert.Equal(t, createdManager.UserID, manager.UserID)
	assert.Equal(t, createdManager.OrganizationID, manager.OrganizationID)
	assert.Equal(t, createdManager.Role, manager.Role)
	assert.Equal(t, createdManager.Name, manager.Name)
}
