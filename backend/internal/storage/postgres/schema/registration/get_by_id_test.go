package registration

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRegistrationByID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	registrationID := uuid.MustParse("80000000-0000-0000-0000-000000000001")

	getInput := &models.GetRegistrationByIDInput{
		ID: registrationID,
	}

	retrieved, err := repo.GetRegistrationByID(ctx, getInput)

	require.Nil(t, err)
	require.NotNil(t, retrieved)
	assert.Equal(t, registrationID, retrieved.Body.ID)
	assert.NotEqual(t, uuid.Nil, retrieved.Body.ChildID)
	assert.NotEqual(t, uuid.Nil, retrieved.Body.GuardianID)
	assert.NotEqual(t, uuid.Nil, retrieved.Body.EventOccurrenceID)
	assert.NotEmpty(t, retrieved.Body.Status)
	assert.NotEmpty(t, retrieved.Body.EventName)
	assert.NotZero(t, retrieved.Body.CreatedAt)
	assert.NotZero(t, retrieved.Body.UpdatedAt)
	assert.NotZero(t, retrieved.Body.OccurrenceStartTime)
}

func TestGetRegistrationByID_VerifyEventDetails(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	registrationID := uuid.MustParse("80000000-0000-0000-0000-000000000006")

	getInput := &models.GetRegistrationByIDInput{
		ID: registrationID,
	}

	retrieved, err := repo.GetRegistrationByID(ctx, getInput)

	require.Nil(t, err)
	require.NotNil(t, retrieved)
	assert.NotEmpty(t, retrieved.Body.EventName)
	assert.NotZero(t, retrieved.Body.OccurrenceStartTime)
}

func TestGetRegistrationByID_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	getInput := &models.GetRegistrationByIDInput{
		ID: uuid.New(),
	}

	retrieved, err := repo.GetRegistrationByID(ctx, getInput)

	require.NotNil(t, err)
	require.Nil(t, retrieved)
}

func TestGetRegistrationByID_MultipleDifferentRegistrations(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	registrationID1 := uuid.MustParse("80000000-0000-0000-0000-000000000001")
	registrationID2 := uuid.MustParse("80000000-0000-0000-0000-000000000002")

	getInput1 := &models.GetRegistrationByIDInput{
		ID: registrationID1,
	}
	retrieved1, err := repo.GetRegistrationByID(ctx, getInput1)
	require.Nil(t, err)
	require.NotNil(t, retrieved1)

	getInput2 := &models.GetRegistrationByIDInput{
		ID: registrationID2,
	}
	retrieved2, err := repo.GetRegistrationByID(ctx, getInput2)
	require.Nil(t, err)
	require.NotNil(t, retrieved2)

	assert.NotEqual(t, retrieved1.Body.ID, retrieved2.Body.ID)
}
