package testutil

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

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
	dbName := fmt.Sprintf("test_%d", time.Now().UnixNano())

	// Create a fresh test DB from the template
	_, err := adminPool.Exec(ctx,
		fmt.Sprintf(`CREATE DATABASE %s TEMPLATE %s`, dbName, templateDB),
	)
	if err != nil {
		t.Fatalf("failed to create test db: %v", err)
	}

	// === NEW: parse the URL and override database ===
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
		_, _ = adminPool.Exec(ctx, fmt.Sprintf(`DROP DATABASE IF EXISTS %s`, dbName))
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
