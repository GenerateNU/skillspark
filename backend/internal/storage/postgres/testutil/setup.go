package testutil

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var cachedConnString string

func ensureContainerAndTemplate() error {
	setupOnce.Do(func() {
		ctx := context.Background()

		log.Println("Starting reusable PostgreSQL container...")

		pg, err := postgres.Run(ctx,
			"postgres:15-alpine",
			postgres.WithDatabase("postgres"),
			postgres.WithUsername("test"),
			postgres.WithPassword("test"),
			testcontainers.WithWaitStrategy(
				wait.ForLog("database system is ready to accept connections").
					WithOccurrence(2),
			),
		)
		if err != nil {
			setupErr = err
			return
		}

		container = pg

		connStr, err := pg.ConnectionString(ctx, "sslmode=disable")
		if err != nil {
			setupErr = err
			return
		}

		cachedConnString = connStr

		adminPool, err = pgxpool.New(ctx, connStr)
		if err != nil {
			setupErr = err
			return
		}

		if err := createTemplateDB(ctx); err != nil {
			setupErr = err
			return
		}

		log.Println("Postgres template DB ready")
	})

	return setupErr
}

func baseConnString() string {
	return cachedConnString
}

func createTemplateDB(ctx context.Context) error {
	_, _ = adminPool.Exec(ctx, fmt.Sprintf(`DROP DATABASE IF EXISTS %s`, templateDB))

	_, err := adminPool.Exec(ctx, fmt.Sprintf(`CREATE DATABASE %s`, templateDB))
	if err != nil {
		return err
	}

	// Parse base URL and override database
	config, err := pgxpool.ParseConfig(baseConnString())
	if err != nil {
		return err
	}
	config.ConnConfig.Database = templateDB

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return err
	}
	defer pool.Close()

	applyFastSettings(pool)

	if err := runSQLDir(pool, "internal/supabase/migrations"); err != nil {
		return fmt.Errorf("migrations failed: %w", err)
	}

	if err := runSQLDir(pool, "internal/supabase/seed"); err != nil {
		return fmt.Errorf("seeds failed: %w", err)
	}

	return nil
}

func applyFastSettings(pool *pgxpool.Pool) {
	ctx := context.Background()
	_, _ = pool.Exec(ctx, `
		SET fsync = off;
		SET synchronous_commit = off;
		SET full_page_writes = off;
	`)
}

func runSQLDir(pool *pgxpool.Pool, dir string) error {
	root, err := findProjectRoot()
	if err != nil {
		return err
	}

	path := filepath.Join(root, dir)
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	var sqlFiles []string
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".sql") {
			sqlFiles = append(sqlFiles, f.Name())
		}
	}
	sort.Strings(sqlFiles)

	for _, name := range sqlFiles {
		content, err := os.ReadFile(filepath.Join(path, name))
		if err != nil {
			return err
		}
		if _, err := pool.Exec(context.Background(), string(content)); err != nil {
			return fmt.Errorf("%s failed: %w", name, err)
		}
	}

	return nil
}

func findProjectRoot() (string, error) {
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("project root not found")
		}
		dir = parent
	}
}
func GetProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("could not find project root (no go.mod found)")
		}
		dir = parent
	}
}
