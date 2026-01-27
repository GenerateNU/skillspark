package manager

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

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
	ptr := uuid.MustParse("40000000-0000-0000-0000-000000000001")
	managerInput := func() *models.PatchManagerInput {
		input := &models.PatchManagerInput{}
		input.Body.ID = uuid.MustParse("50000000-0000-0000-0000-000000000001")
		input.Body.Name = "Updated Assistant"
		input.Body.Email = "updated.assist@example.com"
		input.Body.Username = "uassist"
		input.Body.LanguagePreference = "en"
		input.Body.OrganizationID = &ptr
		input.Body.Role = "Assistant Director"
		return input
	}()

	manager, err := repo.PatchManager(ctx, managerInput)
	assert.Nil(t, err)
	assert.NotNil(t, manager.UserID)
	assert.Equal(t, uuid.MustParse("40000000-0000-0000-0000-000000000001"), manager.OrganizationID)
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
