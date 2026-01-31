package manager

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestManagerRepository_GetManagerByUserID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewManagerRepository(testDB)
	ctx := context.Background()

	orgID := uuid.MustParse("40000000-0000-0000-0000-000000000006")

	input := &models.CreateManagerInput{}
	input.Body.Name = "User ID Manager"
	input.Body.Email = "userid.mgr@test.com"
	input.Body.Username = "uidmgr"
	input.Body.LanguagePreference = "en"
	input.Body.OrganizationID = &orgID
	input.Body.Role = "Admin"

	createdManager, err := repo.CreateManager(ctx, input)
	assert.Nil(t, err)

	manager, err := repo.GetManagerByUserID(ctx, createdManager.UserID)

	assert.Nil(t, err)
	assert.NotNil(t, manager)
	assert.Equal(t, createdManager.ID, manager.ID)
	assert.Equal(t, createdManager.UserID, manager.UserID)
	assert.Equal(t, "Admin", manager.Role)
	assert.Equal(t, "User ID Manager", manager.Name)

	managerTwo, err := repo.GetManagerByUserID(ctx, uuid.Nil)

	assert.NotNil(t, err)
	assert.Nil(t, managerTwo)
}
