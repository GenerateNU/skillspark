package schema

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"skillspark/internal/storage/postgres/testutil"
	"strings"
)

// Reads a SQL file from the schema directory **only from the following directory: 	/backend/storage/postgres/schema**
func ReadSQLBaseScript(path string) (string, error) {
	if !strings.HasSuffix(path, ".sql") {
		return "", errors.New("file is not a SQL file: " + path)
	}

	projectRoot, err := testutil.GetProjectRoot()
	if err != nil {
		return "", fmt.Errorf("failed to get project root: %w", err)
	}

	path = filepath.Join(projectRoot, "internal", "storage", "postgres", "schema", path)

	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
