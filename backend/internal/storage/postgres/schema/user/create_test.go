package user

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create_User(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewUserRepository(testDB)
	ctx := context.Background()

	userInput := func() *models.CreateUserInput {
		input := &models.CreateUserInput{}
		input.Body.Name = "Test User"
		input.Body.Email = "testuser@example.com"
		input.Body.Username = "testuser"
		input.Body.LanguagePreference = "en"
		authID := uuid.New()
		input.Body.AuthID = &authID
		return input
	}()

	user, err := repo.CreateUser(ctx, userInput)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.NotNil(t, user.ID)
	assert.NotNil(t, user.CreatedAt)
	assert.NotNil(t, user.UpdatedAt)
	assert.Equal(t, userInput.Body.Name, user.Name)
	assert.Equal(t, userInput.Body.Email, user.Email)
	assert.Equal(t, userInput.Body.Username, user.Username)
	assert.Equal(t, userInput.Body.AuthID, user.AuthID)

	retrievedUser, err := repo.GetUserByID(ctx, user.ID)
	if err != nil {
		t.Fatalf("Failed to retrieve user: %v", err)
	}
	assert.NotNil(t, retrievedUser)
	assert.Equal(t, user.ID, retrievedUser.ID)
	assert.Equal(t, userInput.Body.Name, retrievedUser.Name)
}

func TestUserRepository_Create_Constraints(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewUserRepository(testDB)
	ctx := context.Background()

	t.Run("Duplicate Email Failure", func(t *testing.T) {
		input1 := &models.CreateUserInput{}
		input1.Body.Name = "User One"
		input1.Body.Email = "duplicate@test.com"
		input1.Body.Username = "userone"
		authID1 := uuid.New()
		input1.Body.AuthID = &authID1

		_, err := repo.CreateUser(ctx, input1)
		assert.NoError(t, err)

		input2 := &models.CreateUserInput{}
		input2.Body.Name = "User Two"
		input2.Body.Email = "duplicate@test.com"
		input2.Body.Username = "usertwo"
		authID2 := uuid.New()
		input2.Body.AuthID = &authID2

		user, err := repo.CreateUser(ctx, input2)
		assert.Error(t, err)
		assert.Nil(t, user)
	})

	t.Run("Duplicate AuthID Failure", func(t *testing.T) {
		input1 := &models.CreateUserInput{}
		input1.Body.Name = "User A"
		input1.Body.Email = "uniqueA@test.com"
		input1.Body.Username = "usera"
		authID := uuid.New()
		input1.Body.AuthID = &authID

		_, err := repo.CreateUser(ctx, input1)
		assert.NoError(t, err)

		input2 := &models.CreateUserInput{}
		input2.Body.Name = "User B"
		input2.Body.Email = "uniqueB@test.com"
		input2.Body.Username = "userb"
		input2.Body.AuthID = input1.Body.AuthID

		user, err := repo.CreateUser(ctx, input2)
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}
