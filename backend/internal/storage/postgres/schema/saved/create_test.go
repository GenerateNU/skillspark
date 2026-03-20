package saved

import (
	"context"
	"skillspark/internal/models"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"
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

	event_occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)
	guardian := guardian.CreateTestGuardian(t, ctx, testDB)
	eventOccurrenceID := event_occurrence.ID
	guardianID := guardian.ID

	input := func() *models.CreateSavedInput {
		i := &models.CreateSavedInput{}
		i.Body.EventOccurrenceID = eventOccurrenceID
		i.Body.GuardianID = guardianID
		return i
	}()

	created, err := repo.CreateSaved(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, created)

	assert.Equal(t, eventOccurrenceID, created.EventOccurrenceID)
	assert.Equal(t, guardianID, created.GuardianID)
	assert.NotEqual(t, uuid.Nil, created.ID)
	assert.NotZero(t, created.CreatedAt)
	assert.NotZero(t, created.UpdatedAt)
}

func TestCreateSaved_FailsInvalidEventOccurrence(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewSavedRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardian := guardian.CreateTestGuardian(t, ctx, testDB)
	guardianID := guardian.ID

	input := func() *models.CreateSavedInput {
		i := &models.CreateSavedInput{}
		i.Body.EventOccurrenceID = uuid.New()
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

	event_occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)
	eventOccurrenceID := event_occurrence.ID

	input := func() *models.CreateSavedInput {
		i := &models.CreateSavedInput{}
		i.Body.EventOccurrenceID = eventOccurrenceID
		i.Body.GuardianID = uuid.New()
		return i
	}()

	created, err := repo.CreateSaved(ctx, input)

	require.NotNil(t, err)
	assert.Nil(t, created)
	assert.Contains(t, err.Error(), "foreign key")
}
