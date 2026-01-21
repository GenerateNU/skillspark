package testutil

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
)

var (
	setupOnce sync.Once
	setupErr  error

	adminPool  *pgxpool.Pool
	container  testcontainers.Container
	templateDB = "template_test_db"
)

// ---------- PUBLIC API ----------

// SetupTestDB returns a pool connected to a fresh database cloned
// from the template DB. Safe for t.Parallel().
func SetupTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()
	t.Parallel()

	if err := ensureContainerAndTemplate(); err != nil {
		t.Fatalf("test DB setup failed: %v", err)
	}

	ctx := context.Background()
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	dbName := fmt.Sprintf("test_%d_%d", time.Now().UnixNano(), rng.Intn(1_000_000))

	dbIdent := pgx.Identifier{dbName}.Sanitize()
	templateIdent := pgx.Identifier{templateDB}.Sanitize()

	// Create a fresh test DB from the template
	_, err := adminPool.Exec(
		ctx,
		fmt.Sprintf(`CREATE DATABASE %s TEMPLATE %s`, dbIdent, templateIdent),
	)
	if err != nil {
		t.Fatalf("failed to create test db: %v", err)
	}

	// Parse base connection string and override database
	config, err := pgxpool.ParseConfig(baseConnString())
	if err != nil {
		t.Fatalf("failed to parse base connection string: %v", err)
	}
	config.ConnConfig.Database = dbName

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		t.Fatalf("failed to connect test db: %v", err)
	}

	t.Cleanup(func() {
		pool.Close()
		_, _ = adminPool.Exec(
			ctx,
			fmt.Sprintf(`DROP DATABASE IF EXISTS %s`, dbIdent),
		)
	})

	return pool
}

// Shutdown terminates the shared container (call from TestMain)
func Shutdown() {
	if adminPool != nil {
		adminPool.Close()
	}
	if container != nil {
		_ = container.Terminate(context.Background())
	}
}
