package guardian

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
)

func CreateTestGuardian(
	t *testing.T,
	ctx context.Context,
) *models.Guardian {
	t.Helper()

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)

	input := &models.CreateGuardianInput{}

	// get a user
	input.Body.UserID = uuid.MustParse("f2a3b4c5-d6e7-4f8a-9b0c-1d2e3f4a5b6c")

	guardian, _ := repo.CreateGuardian(ctx, input)

	return guardian
}
