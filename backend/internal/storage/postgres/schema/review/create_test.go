package review

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/registration"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateReview(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewReviewRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	registration := registration.CreateTestRegistration(t, ctx, testDB)
	guardianID := registration.GuardianID
	registrationID := registration.ID

	input := func() *models.CreateReviewInput {
		i := &models.CreateReviewInput{}
		i.Body.RegistrationID = registrationID
		i.Body.GuardianID = guardianID
		i.Body.Description = "Test review"
		i.Body.Categories = []string{"informative"}
		return i
	}()

	created, err := repo.CreateReview(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, created)

	assert.Equal(t, registrationID, created.RegistrationID)
	assert.Equal(t, guardianID, created.GuardianID)
	assert.Equal(t, input.Body.Description, created.Description)
	assert.Equal(t, input.Body.Categories, created.Categories)
	assert.NotEqual(t, uuid.Nil, created.ID)
	assert.NotZero(t, created.CreatedAt)
	assert.NotZero(t, created.UpdatedAt)
}

func TestCreateReview_FailsInvalidRegistration(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewReviewRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	validRegistration := registration.CreateTestRegistration(t, ctx, testDB)
	validGuardianID := validRegistration.GuardianID

	input := func() *models.CreateReviewInput {
		i := &models.CreateReviewInput{}
		i.Body.RegistrationID = uuid.New()
		i.Body.GuardianID = validGuardianID
		i.Body.Description = "Test review"
		i.Body.Categories = []string{"informative"}
		return i
	}()

	created, err := repo.CreateReview(ctx, input)

	require.NotNil(t, err)
	assert.Nil(t, created)
	assert.Contains(t, err.Error(), "foreign key")
}

func TestCreateReview_FailsInvalidGuardian(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewReviewRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	registration := registration.CreateTestRegistration(t, ctx, testDB)

	input := func() *models.CreateReviewInput {
		i := &models.CreateReviewInput{}
		i.Body.RegistrationID = registration.ID
		i.Body.GuardianID = uuid.New()
		i.Body.Description = "Test review"
		i.Body.Categories = []string{"informative"}
		return i
	}()

	created, err := repo.CreateReview(ctx, input)

	require.NotNil(t, err)
	assert.Nil(t, created)
	assert.Contains(t, err.Error(), "foreign key")
}