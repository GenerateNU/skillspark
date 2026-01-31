package manager

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/organization"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
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
		input.Body.UserID = uuid.MustParse("f6a7b8c9-d0e1-4f2a-3b4c-5d6e7f8a9b0c")
		input.Body.OrganizationID = &organizationID
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

}
