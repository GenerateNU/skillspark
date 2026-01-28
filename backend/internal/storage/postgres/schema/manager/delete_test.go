package manager

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// -------------------- Director --------------------
func TestManagerRepository_Delete_Director(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewManagerRepository(testDB)
	ctx := context.Background()
	ptr := uuid.MustParse("40000000-0000-0000-0000-000000000006")
	managerInput := func() *models.CreateManagerInput {
		input := &models.CreateManagerInput{}
		input.Body.Name = "Delete Man"
		input.Body.Email = "delete.m@example.com"
		input.Body.Username = "delman"
		input.Body.LanguagePreference = "en"
		input.Body.OrganizationID = &ptr
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
