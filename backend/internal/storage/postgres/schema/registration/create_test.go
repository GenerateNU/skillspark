package registration

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/child"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateRegistration(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	child := child.CreateTestChild(t, ctx, testDB)
	occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)
	childID := child.ID
	guardianID := child.GuardianID
	occurrenceID := occurrence.ID

	input := func() *models.CreateRegistrationInput {
		i := &models.CreateRegistrationInput{}
		i.Body.ChildID = child.ID
		i.Body.GuardianID = child.GuardianID
		i.Body.EventOccurrenceID = occurrence.ID
		i.Body.Status = models.RegistrationStatusRegistered
		return i
	}()

	created, err := repo.CreateRegistration(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, created)
	assert.Equal(t, &childID, created.Body.ChildID)
	assert.Equal(t, &guardianID, created.Body.GuardianID)
	assert.Equal(t, occurrenceID, created.Body.EventOccurrenceID)
	assert.Equal(t, models.RegistrationStatusRegistered, created.Body.Status)
	assert.NotEqual(t, uuid.Nil, created.Body.ID)
	assert.NotZero(t, created.Body.CreatedAt)
	assert.NotZero(t, created.Body.UpdatedAt)
	assert.NotEmpty(t, created.Body.EventName)
	assert.NotZero(t, created.Body.OccurrenceStartTime)
}

func TestCreateRegistration_VerifyEventNameJoin(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	child := child.CreateTestChild(t, ctx, testDB)
	occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)

	input := func() *models.CreateRegistrationInput {
		i := &models.CreateRegistrationInput{}
		i.Body.ChildID = child.ID
		i.Body.GuardianID = child.GuardianID
		i.Body.EventOccurrenceID = occurrence.ID
		i.Body.Status = models.RegistrationStatusRegistered
		return i
	}()

	created, err := repo.CreateRegistration(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, created)
	assert.NotEmpty(t, created.Body.EventName)
}

func TestCreateRegistration_VerifyOccurrenceStartTime(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	child := child.CreateTestChild(t, ctx, testDB)
	occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)
	childID := child.ID
	guardianID := child.GuardianID
	occurrenceID := occurrence.ID

	input := func() *models.CreateRegistrationInput {
		i := &models.CreateRegistrationInput{}
		i.Body.ChildID = childID
		i.Body.GuardianID = guardianID
		i.Body.EventOccurrenceID = occurrenceID
		i.Body.Status = models.RegistrationStatusRegistered
		return i
	}()

	created, err := repo.CreateRegistration(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, created)
	assert.NotZero(t, created.Body.OccurrenceStartTime)
}

func TestCreateRegistration_MultipleRegistrationsForSameChild(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	child := child.CreateTestChild(t, ctx, testDB)
	childID := child.ID
	guardianID := child.GuardianID
	o1 := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)
	o1ID := o1.ID
	o2 := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)
	o2ID := o2.ID

	input1 := func() *models.CreateRegistrationInput {
		i := &models.CreateRegistrationInput{}
		i.Body.ChildID = childID
		i.Body.GuardianID = guardianID
		i.Body.EventOccurrenceID = o1ID
		i.Body.Status = models.RegistrationStatusRegistered
		return i
	}()

	created1, err := repo.CreateRegistration(ctx, input1)
	require.Nil(t, err)
	require.NotNil(t, created1)

	input2 := func() *models.CreateRegistrationInput {
		i := &models.CreateRegistrationInput{}
		i.Body.ChildID = childID
		i.Body.GuardianID = guardianID
		i.Body.EventOccurrenceID = o2ID
		i.Body.Status = models.RegistrationStatusRegistered
		return i
	}()

	created2, err := repo.CreateRegistration(ctx, input2)
	require.Nil(t, err)
	require.NotNil(t, created2)

	assert.NotEqual(t, created1.Body.ID, created2.Body.ID)
	assert.NotEqual(t, created1.Body.EventOccurrenceID, created2.Body.EventOccurrenceID)
}

func TestCreateRegistration_InvalidChildID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	child := child.CreateTestChild(t, ctx, testDB)
	occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)
	guardianID := child.GuardianID
	occurrenceID := occurrence.ID

	input := func() *models.CreateRegistrationInput {
		i := &models.CreateRegistrationInput{}
		i.Body.ChildID = uuid.New()
		i.Body.GuardianID = guardianID
		i.Body.EventOccurrenceID = occurrenceID
		i.Body.Status = models.RegistrationStatusRegistered
		return i
	}()

	created, err := repo.CreateRegistration(ctx, input)

	require.NotNil(t, err)
	require.Nil(t, created)
}

func TestCreateRegistration_InvalidGuardianID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	child := child.CreateTestChild(t, ctx, testDB)
	occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, testDB)
	childID := child.ID
	occurrenceID := occurrence.ID

	input := func() *models.CreateRegistrationInput {
		i := &models.CreateRegistrationInput{}
		i.Body.ChildID = childID
		i.Body.GuardianID = uuid.New()
		i.Body.EventOccurrenceID = occurrenceID
		i.Body.Status = models.RegistrationStatusRegistered
		return i
	}()

	created, err := repo.CreateRegistration(ctx, input)

	require.NotNil(t, err)
	require.Nil(t, created)
}

func TestCreateRegistration_InvalidEventOccurrenceID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	child := child.CreateTestChild(t, ctx, testDB)
	childID := child.ID
	guardianID := child.GuardianID

	input := func() *models.CreateRegistrationInput {
		i := &models.CreateRegistrationInput{}
		i.Body.ChildID = childID
		i.Body.GuardianID = guardianID
		i.Body.EventOccurrenceID = uuid.New()
		i.Body.Status = models.RegistrationStatusRegistered
		return i
	}()

	created, err := repo.CreateRegistration(ctx, input)

	require.NotNil(t, err)
	require.Nil(t, created)
}
