package guardian

import (
	"context"
	"skillspark/internal/models"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func CreateTestGuardian(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Guardian {
	t.Helper()

	repo := NewGuardianRepository(db)

	input := &models.CreateGuardianInput{}

	// get a user
	input.Body.UserID = uuid.MustParse("f2a3b4c5-d6e7-4f8a-9b0c-1d2e3f4a5b6c")

	guardian, err := repo.CreateGuardian(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, guardian)

	return guardian
}
