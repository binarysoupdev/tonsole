package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// Create a new test directory and it join with path.
func NewPath(t *testing.T, path string) string {
	return filepath.Join(t.TempDir(), path)
}

// Create an empty file in a new test directory and return its path.
func CreateEmpty(t *testing.T, path string) string {
	file, path := Create(t, path)
	file.Close()

	return path
}

// Create a new file in a new test directory and return the open file pointer and its path.
func Create(t *testing.T, path string) (*os.File, string) {
	path = NewPath(t, path)

	file, err := os.Create(path)
	require.NoError(t, err)

	return file, path
}
