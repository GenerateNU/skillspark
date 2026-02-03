package manager

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestManagerRepository_GetManagerByOrgID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewManagerRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	orgID := uuid.MustParse("40000000-0000-0000-0000-000000000006")

	input := &models.CreateManagerInput{}
	input.Body.Name = "Org Manager"
	input.Body.Email = "org.mgr@test.com"
	input.Body.Username = "orgmgr"
	input.Body.LanguagePreference = "en"
	input.Body.OrganizationID = &orgID
	input.Body.Role = "Director"

	createdManager, err := repo.CreateManager(ctx, input)
	assert.Nil(t, err)

	manager, err := repo.GetManagerByOrgID(ctx, orgID)

	assert.Nil(t, err)
	assert.NotNil(t, manager)
	assert.Equal(t, createdManager.UserID, manager.UserID)
	assert.Equal(t, createdManager.OrganizationID, manager.OrganizationID)
	assert.Equal(t, "Director", manager.Role)
	assert.Equal(t, "Org Manager", manager.Name)

	managerThree, err := repo.GetManagerByOrgID(ctx, uuid.MustParse("00000000-0000-0000-0000-000000000000"))

	assert.NotNil(t, err)
	assert.Nil(t, managerThree)
}
