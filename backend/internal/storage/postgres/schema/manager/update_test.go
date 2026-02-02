package manager

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/organization"
	"skillspark/internal/storage/postgres/testutil"
	"skillspark/internal/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// -------------------- Assistant Director --------------------
func TestManagerRepository_Update_AssistantDirector(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewManagerRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	organizationID := organization.CreateTestOrganization(t, ctx, testDB).ID
	managerInput := func() *models.PatchManagerInput {
		input := &models.PatchManagerInput{}
		input.Body.ID = uuid.MustParse("50000000-0000-0000-0000-000000000001")
		input.Body.OrganizationID = &organizationID
		input.Body.Name = utils.PtrString("Updated Assistant")
		input.Body.Email = utils.PtrString("updated.assist@example.com")
		input.Body.Username = utils.PtrString("uassist")
		input.Body.LanguagePreference = utils.PtrString("en")
		input.Body.Role = utils.PtrString("Assistant Director")
		return input
	}()

	manager, err := repo.PatchManager(ctx, managerInput)
	assert.Nil(t, err)
	assert.NotNil(t, manager.UserID)
	assert.Equal(t, organizationID, manager.OrganizationID)
	assert.Equal(t, "Assistant Director", manager.Role)
	assert.Equal(t, "Updated Assistant", manager.Name)

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
