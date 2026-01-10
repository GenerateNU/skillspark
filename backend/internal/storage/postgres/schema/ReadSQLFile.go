package schema

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// Reads a SQL file from the schema directory **only from the following directory: 	/backend/storage/postgres/schema**
func ReadSQLBaseScript(path string) (string, error) {
	if !strings.HasSuffix(path, ".sql") {
		return "", errors.New("file is not a SQL file: " + path)
	}

	currdir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	projectRoot := strings.Split(currdir, "/skillspark")[0] + "/skillspark"

	path = filepath.Join(projectRoot, "backend", "internal", "storage", "postgres", "schema", path)

	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
