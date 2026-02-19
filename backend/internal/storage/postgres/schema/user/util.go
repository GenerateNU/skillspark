package user

import (
	"context"
	"embed"
	"skillspark/internal/models"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

//go:embed sql/*.sql
var SqlUserFiles embed.FS

func CreateTestUser(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.User {
	t.Helper()

	repo := NewUserRepository(db)

	input := &models.CreateUserInput{}
	input.Body.Name = "Test User"
	input.Body.Email = "testuser@example.com"
	input.Body.Username = "testuser"
	input.Body.LanguagePreference = "en"
	authID := uuid.New()
	input.Body.AuthID = authID
	user, err := repo.CreateUser(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, user)

	return user
}
