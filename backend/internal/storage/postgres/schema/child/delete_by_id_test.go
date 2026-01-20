package child

import (
	"context"
	"testing"

	"skillspark/internal/storage/postgres/testutil"

	"github.com/stretchr/testify/assert"
)

func TestChildRepository_DeleteChildByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()

	createdChild := CreateTestChild(t, ctx, repo)

	child, err := repo.DeleteChildByID(ctx, createdChild.ID)

	assert.Nil(t, err)
	assert.NotNil(t, child)

	child, err = repo.GetChildByID(ctx, createdChild.ID)

	assert.Nil(t, child)
	assert.NotNil(t, err)

}
