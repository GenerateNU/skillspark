// testutil/shared_db.go
package testutil

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	// Global shared test database - initialized once, used by all tests
	sharedTestDB *SharedTestDB
	setupOnce    sync.Once
	setupError   error
)

// SharedTestDB holds the shared test database resources
type SharedTestDB struct {
	Pool      *pgxpool.Pool
	container testcontainers.Container
	mu        sync.Mutex
}

// GetSharedTestDB returns the global shared test database
// This function is thread-safe and ensures the container is only started once
func GetSharedTestDB() (*SharedTestDB, error) {
	setupOnce.Do(func() {
		ctx := context.Background()

		log.Println("Starting shared PostgreSQL container...")

		// Start PostgreSQL container with optimizations
		pgContainer, err := postgres.Run(ctx,
			"postgres:15-alpine",
			postgres.WithDatabase("testdb"),
			postgres.WithUsername("test"),
			postgres.WithPassword("test"),
			testcontainers.WithWaitStrategy(
				wait.ForLog("database system is ready to accept connections").
					WithOccurrence(2).
					WithStartupTimeout(30*time.Second),
			),
			// Add test optimizations
			testcontainers.WithEnv(map[string]string{
				"POSTGRES_INITDB_ARGS":      "-E UTF8 --auth-local=trust",
				"POSTGRES_HOST_AUTH_METHOD": "trust",
			}),
		)
		if err != nil {
			setupError = fmt.Errorf("failed to start container: %w", err)
			return
		}

		// Get connection string
		connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
		if err != nil {
			setupError = fmt.Errorf("failed to get connection string: %w", err)
			return
		}

		// Configure pool for testing
		config, err := pgxpool.ParseConfig(connStr)
		if err != nil {
			setupError = fmt.Errorf("failed to parse config: %w", err)
			return
		}

		// Optimize pool for testing
		config.MaxConns = 50 // Allow many parallel tests
		config.MinConns = 10 // Keep connections warm
		config.MaxConnLifetime = time.Hour
		config.MaxConnIdleTime = time.Minute * 30

		// Create pool
		pool, err := pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			setupError = fmt.Errorf("failed to create pool: %w", err)
			return
		}

		sharedTestDB = &SharedTestDB{
			Pool:      pool,
			container: pgContainer,
		}

		// Apply test optimizations to PostgreSQL
		applyTestOptimizations(pool)

		// Run migrations to create all tables
		if err := runMigrations(pool); err != nil {
			setupError = fmt.Errorf("failed to run migrations: %w", err)
			return
		}

		log.Println("Shared test database ready!")
	})

	return sharedTestDB, setupError
}

// applyTestOptimizations applies PostgreSQL settings for faster tests
func applyTestOptimizations(pool *pgxpool.Pool) {
	ctx := context.Background()

	// ONLY for tests - makes writes much faster but less safe
	optimizations := []string{
		"ALTER SYSTEM SET fsync = off",
		"ALTER SYSTEM SET synchronous_commit = off",
		"ALTER SYSTEM SET full_page_writes = off",
		"ALTER SYSTEM SET checkpoint_segments = 100",
		"ALTER SYSTEM SET checkpoint_completion_target = 0.9",
		"ALTER SYSTEM SET wal_buffers = '64MB'",
		"ALTER SYSTEM SET shared_buffers = '256MB'",
		"SELECT pg_reload_conf()",
	}

	for _, sql := range optimizations {
		if _, err := pool.Exec(ctx, sql); err != nil {
			log.Printf("Warning: Failed to apply optimization %q: %v", sql, err)
		}
	}
}

// runMigrations reads and executes all migration files from /supabase/migrations
func runMigrations(pool *pgxpool.Pool) error {
	ctx := context.Background()

	// Find the migrations directory
	migrationsPath, err := findMigrationsPath()
	if err != nil {
		return fmt.Errorf("failed to find migrations directory: %w", err)
	}

	log.Printf("Reading migrations from: %s", migrationsPath)

	// Read all .sql files from the migrations directory
	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	// Filter and sort migration files
	var migrationFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			migrationFiles = append(migrationFiles, file.Name())
		}
	}

	if len(migrationFiles) == 0 {
		return fmt.Errorf("no migration files found in %s", migrationsPath)
	}

	// Sort files to ensure they run in order (assuming timestamp or numbered prefixes)
	sort.Strings(migrationFiles)

	log.Printf("Found %d migration files", len(migrationFiles))

	// Execute each migration file
	for _, filename := range migrationFiles {
		log.Printf("Running migration: %s", filename)

		filePath := filepath.Join(migrationsPath, filename)
		content, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", filename, err)
		}

		// Execute the migration in a transaction
		tx, err := pool.Begin(ctx)
		if err != nil {
			return fmt.Errorf("failed to begin transaction for %s: %w", filename, err)
		}

		if _, err := tx.Exec(ctx, string(content)); err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("failed to execute migration %s: %w", filename, err)
		}

		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("failed to commit migration %s: %w", filename, err)
		}
	}

	log.Println("All migrations completed successfully")
	return nil
}

// findMigrationsPath finds the /supabase/migrations directory
// It searches upward from the current directory until it finds it
func findMigrationsPath() (string, error) {
	// Start from current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Search up the directory tree
	currentPath := cwd
	for {
		migrationsPath := filepath.Join(currentPath, "supabase", "migrations")
		if _, err := os.Stat(migrationsPath); err == nil {
			return migrationsPath, nil
		}

		// Move up one directory
		parentPath := filepath.Dir(currentPath)
		if parentPath == currentPath {
			// Reached root without finding migrations
			break
		}
		currentPath = parentPath
	}

	return "", fmt.Errorf("could not find /supabase/migrations directory (searched from %s)", cwd)
}

// CleanupTestData truncates all tables efficiently
// This is much faster than DELETE and resets auto-increment counters
func (db *SharedTestDB) CleanupTestData(t testing.TB) {
	t.Helper()

	// Ensure cleanup is thread-safe
	db.mu.Lock()
	defer db.mu.Unlock()

	ctx := context.Background()

	// Get all table names from the database
	rows, err := db.Pool.Query(ctx, `
		SELECT tablename 
		FROM pg_tables 
		WHERE schemaname = 'public'
		ORDER BY tablename
	`)
	if err != nil {
		t.Fatalf("Failed to get table names: %v", err)
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			t.Fatalf("Failed to scan table name: %v", err)
		}
		tables = append(tables, tableName)
	}

	if err := rows.Err(); err != nil {
		t.Fatalf("Error iterating table names: %v", err)
	}

	if len(tables) == 0 {
		return // No tables to clean
	}

	// Build TRUNCATE statement for all tables
	// RESTART IDENTITY resets sequences
	// CASCADE handles foreign key dependencies
	truncateSQL := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE",
		strings.Join(tables, ", "))

	if _, err := db.Pool.Exec(ctx, truncateSQL); err != nil {
		t.Fatalf("Failed to cleanup test data: %v", err)
	}
}

// Shutdown closes the shared test database
// Call this in TestMain after all tests complete
func Shutdown() {
	if sharedTestDB != nil {
		sharedTestDB.Pool.Close()

		if sharedTestDB.container != nil {
			ctx := context.Background()
			if err := sharedTestDB.container.Terminate(ctx); err != nil {
				log.Printf("Warning: Failed to terminate container: %v", err)
			}
		}
	}
}

// SetupTestWithCleanup is a helper that gets the shared DB and ensures cleanup
func SetupTestWithCleanup(t testing.TB) *pgxpool.Pool {
	t.Helper()

	db, err := GetSharedTestDB()
	if err != nil {
		t.Fatalf("Failed to get shared test database: %v", err)
	}

	// Clean before test (not after, so failed tests leave data for debugging)
	db.CleanupTestData(t)

	return db.Pool
}
