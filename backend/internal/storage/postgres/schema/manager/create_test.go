package manager

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/organization"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/stretchr/testify/assert"
)

func TestManagerRepository_Create_AssistantDirector(t *testing.T) {
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
		input.Body.Name = "Assistant Man"
		input.Body.Email = "am@example.com"
		input.Body.Username = "amanager"
		input.Body.LanguagePreference = "en"
		input.Body.Role = "Assistant Manager"
		return input
	}()

	manager, err := repo.CreateManager(ctx, managerInput)
	assert.Nil(t, err)
	assert.NotNil(t, manager.UserID)
	assert.Equal(t, organizationID, manager.OrganizationID)
	assert.Equal(t, "Assistant Manager", manager.Role)
	assert.Equal(t, "Assistant Man", manager.Name)

	id := manager.ID

	// Verify we can retrieve the created location
	retrievedManager, err := repo.GetManagerByID(ctx, id)
	assert.Nil(t, err)
	assert.NotNil(t, retrievedManager)
	assert.Equal(t, manager.ID, retrievedManager.ID)
	assert.Equal(t, manager.UserID, retrievedManager.UserID)
	assert.Equal(t, manager.OrganizationID, retrievedManager.OrganizationID)
	assert.Equal(t, manager.Role, retrievedManager.Role)
	assert.Equal(t, manager.Name, retrievedManager.Name)
}
