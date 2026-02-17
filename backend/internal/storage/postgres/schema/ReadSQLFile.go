package schema

import (
	"embed"
	"errors"
	"strings"
)

func ReadSQLBaseScript(path string, fs embed.FS) (string, error) {
	if !strings.HasSuffix(path, ".sql") {
		return "", errors.New("file is not a SQL file: " + path)
	}

	// Prepend "sql/" if not already present, since embed.FS includes the directory prefix
	if !strings.HasPrefix(path, "sql/") {
		path = "sql/" + path
	}

	content, err := fs.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
