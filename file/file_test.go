package file_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/binarysoupdev/tinsel/file"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateFile(t *testing.T) {
	//-- arrange
	const SEED = 42
	r := rand.New(SEED)

	FILE := fmt.Sprintf("%s.%s", r.ASCII(10), r.ASCII(3))

	//-- act
	path := file.CreateEmpty(t, FILE)

	//-- assert
	stat, err := os.Stat(path)

	require.NoError(t, err)
	assert.Equal(t, FILE, stat.Name())
}
