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

func TestDeleteRegistration(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	registrationID := uuid.MustParse("80000000-0000-0000-0000-000000000001")

	getInput := &models.GetRegistrationByIDInput{
		ID: registrationID,
	}
	existing, getErr := repo.GetRegistrationByID(ctx, getInput)
	require.Nil(t, getErr)
	require.NotNil(t, existing)

	deleteInput := &models.DeleteRegistrationInput{
		ID: registrationID,
	}

	deleted, deleteErr := repo.DeleteRegistration(ctx, deleteInput)
	require.Nil(t, deleteErr)
	require.NotNil(t, deleted)
	assert.Equal(t, registrationID, deleted.Body.ID)
	assert.Equal(t, existing.Body.ChildID, deleted.Body.ChildID)
	assert.Equal(t, existing.Body.GuardianID, deleted.Body.GuardianID)
	assert.Equal(t, existing.Body.EventOccurrenceID, deleted.Body.EventOccurrenceID)
	assert.Equal(t, existing.Body.Status, deleted.Body.Status)
	assert.Equal(t, existing.Body.EventName, deleted.Body.EventName)
	assert.NotZero(t, deleted.Body.OccurrenceStartTime)

	_, getErr2 := repo.GetRegistrationByID(ctx, getInput)
	assert.NotNil(t, getErr2)
}

func TestDeleteRegistration_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	deleteInput := &models.DeleteRegistrationInput{
		ID: uuid.New(),
	}

	deleted, err := repo.DeleteRegistration(ctx, deleteInput)

	require.NotNil(t, err)
	assert.Nil(t, deleted)
}

func TestDeleteRegistration_AlreadyDeleted(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	registrationID := uuid.MustParse("80000000-0000-0000-0000-000000000002")

	deleteInput1 := &models.DeleteRegistrationInput{
		ID: registrationID,
	}
	deleted1, deleteErr1 := repo.DeleteRegistration(ctx, deleteInput1)
	require.Nil(t, deleteErr1)
	require.NotNil(t, deleted1)

	deleteInput2 := &models.DeleteRegistrationInput{
		ID: registrationID,
	}
	deleted2, deleteErr2 := repo.DeleteRegistration(ctx, deleteInput2)
	require.NotNil(t, deleteErr2)
	assert.Nil(t, deleted2)
}