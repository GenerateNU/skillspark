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

func TestGetRegistrationsByChildID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	childID := uuid.MustParse("30000000-0000-0000-0000-000000000001")

	input := &models.GetRegistrationsByChildIDInput{
		ChildID: childID,
	}

	result, err := repo.GetRegistrationsByChildID(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, result)
	assert.NotEmpty(t, result.Body.Registrations)
	
	for _, reg := range result.Body.Registrations {
		assert.Equal(t, childID, reg.ChildID)
		assert.NotEqual(t, uuid.Nil, reg.ID)
		assert.NotEqual(t, uuid.Nil, reg.GuardianID)
		assert.NotEqual(t, uuid.Nil, reg.EventOccurrenceID)
		assert.NotEmpty(t, reg.Status)
		assert.NotEmpty(t, reg.EventName)
		assert.NotZero(t, reg.CreatedAt)
		assert.NotZero(t, reg.UpdatedAt)
		assert.NotZero(t, reg.OccurrenceStartTime)
	}
}

func TestGetRegistrationsByChildID_MultipleRegistrations(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	childID := uuid.MustParse("30000000-0000-0000-0000-000000000001")

	input := &models.GetRegistrationsByChildIDInput{
		ChildID: childID,
	}

	result, err := repo.GetRegistrationsByChildID(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, result)
	assert.GreaterOrEqual(t, len(result.Body.Registrations), 2)

	registrationIDs := make(map[uuid.UUID]bool)
	for _, reg := range result.Body.Registrations {
		assert.Equal(t, childID, reg.ChildID)
		registrationIDs[reg.ID] = true
	}
	
	assert.Equal(t, len(result.Body.Registrations), len(registrationIDs))
}

func TestGetRegistrationsByChildID_NoRegistrations(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	childID := uuid.New()

	input := &models.GetRegistrationsByChildIDInput{
		ChildID: childID,
	}

	result, err := repo.GetRegistrationsByChildID(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, result)
	assert.Empty(t, result.Body.Registrations)
}

func TestGetRegistrationsByChildID_VerifyEventDetails(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	ctx := context.Background()

	childID := uuid.MustParse("30000000-0000-0000-0000-000000000003")

	input := &models.GetRegistrationsByChildIDInput{
		ChildID: childID,
	}

	result, err := repo.GetRegistrationsByChildID(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, result)
	assert.NotEmpty(t, result.Body.Registrations)

	for _, reg := range result.Body.Registrations {
		assert.NotEmpty(t, reg.EventName)
		assert.NotZero(t, reg.OccurrenceStartTime)
	}
}