package registration

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateRegistration(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	childID := uuid.MustParse("30000000-0000-0000-0000-000000000001")
	guardianID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	occurrenceID := uuid.MustParse("70000000-0000-0000-0000-000000000002")

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
	assert.Equal(t, childID, created.Body.ChildID)
	assert.Equal(t, guardianID, created.Body.GuardianID)
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

	childID := uuid.MustParse("30000000-0000-0000-0000-000000000002")
	guardianID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	occurrenceID := uuid.MustParse("70000000-0000-0000-0000-000000000001")

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
	assert.NotEmpty(t, created.Body.EventName)
}

func TestCreateRegistration_VerifyOccurrenceStartTime(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	childID := uuid.MustParse("30000000-0000-0000-0000-000000000003")
	guardianID := uuid.MustParse("22222222-2222-2222-2222-222222222222")
	occurrenceID := uuid.MustParse("70000000-0000-0000-0000-000000000005")

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
	assert.True(t, created.Body.OccurrenceStartTime.After(time.Now()))
}

func TestCreateRegistration_MultipleRegistrationsForSameChild(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	childID := uuid.MustParse("30000000-0000-0000-0000-000000000004")
	guardianID := uuid.MustParse("33333333-3333-3333-3333-333333333333")

	input1 := func() *models.CreateRegistrationInput {
		i := &models.CreateRegistrationInput{}
		i.Body.ChildID = childID
		i.Body.GuardianID = guardianID
		i.Body.EventOccurrenceID = uuid.MustParse("70000000-0000-0000-0000-000000000002")
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
		i.Body.EventOccurrenceID = uuid.MustParse("70000000-0000-0000-0000-000000000003")
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

	guardianID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	occurrenceID := uuid.MustParse("70000000-0000-0000-0000-000000000001")

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

	childID := uuid.MustParse("30000000-0000-0000-0000-000000000001")
	occurrenceID := uuid.MustParse("70000000-0000-0000-0000-000000000001")

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

	childID := uuid.MustParse("30000000-0000-0000-0000-000000000001")
	guardianID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

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
