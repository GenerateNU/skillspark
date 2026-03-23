package saved

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/event"
	"skillspark/internal/storage/postgres/schema/guardian"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateSaved(t *testing.T) {

	testDB := testutil.SetupTestDB(t)
	repo := NewSavedRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	event := event.CreateTestEvent(t, ctx, testDB)
	guardian := guardian.CreateTestGuardian(t, ctx, testDB)
	guardianID := guardian.ID

	input := func() *models.CreateSavedInput {
		i := &models.CreateSavedInput{}
		i.Body.EventID = event.ID
		i.Body.GuardianID = guardianID
		return i
	}()

	created, err := repo.CreateSaved(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, created)

	assert.Equal(t, event.ID, created.Event.ID)
	assert.Equal(t, guardianID, created.GuardianID)
	assert.NotEqual(t, uuid.Nil, created.ID)
	assert.NotZero(t, created.CreatedAt)
	assert.NotZero(t, created.UpdatedAt)
}

func TestCreateSaved_FailsInvalidEvent(t *testing.T) {

	testDB := testutil.SetupTestDB(t)
	repo := NewSavedRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardian := guardian.CreateTestGuardian(t, ctx, testDB)
	guardianID := guardian.ID

	input := func() *models.CreateSavedInput {
		i := &models.CreateSavedInput{}
		i.Body.EventID = uuid.New()
		i.Body.GuardianID = guardianID
		return i
	}()

	created, err := repo.CreateSaved(ctx, input)

	require.NotNil(t, err)
	assert.Nil(t, created)
	assert.Contains(t, err.Error(), "foreign key")
}

func TestCreateSaved_FailsInvalidGuardian(t *testing.T) {

	testDB := testutil.SetupTestDB(t)
	repo := NewSavedRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	event := event.CreateTestEvent(t, ctx, testDB)

	input := func() *models.CreateSavedInput {
		i := &models.CreateSavedInput{}
		i.Body.EventID = event.ID
		i.Body.GuardianID = uuid.New()
		return i
	}()

	created, err := repo.CreateSaved(ctx, input)

	require.NotNil(t, err)
	assert.Nil(t, created)
	assert.Contains(t, err.Error(), "foreign key")
}
