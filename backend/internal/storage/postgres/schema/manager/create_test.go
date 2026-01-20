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
func TestManagerRepository_Create_AssistantDirector(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewManagerRepository(testDB)
	ctx := context.Background()

	managerInput := func() *models.CreateManagerInput {
		input := &models.CreateManagerInput{}
		input.Body.UserID = uuid.MustParse("f6a7b8c9-d0e1-4f2a-3b4c-5d6e7f8a9b0c")
		input.Body.OrganizationID = uuid.MustParse("40000000-0000-0000-0000-000000000006")
		input.Body.Role = "Assistant Manager"
		return input
	}()

	manager, err := repo.CreateManager(ctx, managerInput)
	assert.Nil(t, err)
	assert.NotNil(t, manager.UserID)
	assert.Equal(t, uuid.MustParse("40000000-0000-0000-0000-000000000006"), manager.OrganizationID)
	assert.Equal(t, "Assistant Manager", manager.Role)

	id := manager.ID

	// Verify we can retrieve the created location
	retrievedManager, err := repo.GetManagerByID(ctx, id)
	assert.Nil(t, err)
	assert.NotNil(t, retrievedManager)
	assert.Equal(t, manager.ID, retrievedManager.ID)
	assert.Equal(t, manager.UserID, retrievedManager.UserID)
	assert.Equal(t, manager.OrganizationID, retrievedManager.OrganizationID)
	assert.Equal(t, manager.Role, retrievedManager.Role)
}
