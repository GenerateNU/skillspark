package guardian

import (
	"context"
	"embed"
	"skillspark/internal/models"
	"testing"

	"math/rand"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

//go:embed sql/*.sql
var SqlGuardianFiles embed.FS

func CreateTestGuardian(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Guardian {
	t.Helper()

	repo := NewGuardianRepository(db)

	input := &models.CreateGuardianInput{}

	input.Body.Email = RandomString(10)
	input.Body.Username = RandomString(10)
	input.Body.AuthID = uuid.New()

	guardian, err := repo.CreateGuardian(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, guardian)

	return guardian
}

func RandomString(n int) string {

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
