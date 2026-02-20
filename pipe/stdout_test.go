package pipe_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/binarysoupdev/tinsel/pipe"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/assert"
)

func TestStdoutPipePrintOnce(t *testing.T) {
	//-- arrange
	const SEED = 42
	r := rand.New(SEED)

	OUTPUT := []string{r.ASCII(15), r.ASCII(15), r.ASCII(15)}

	out := pipe.OpenStdout(len(OUTPUT))
	defer out.Close()

	//-- act
	fmt.Println(strings.Join(OUTPUT, "\n"))

	//-- assert
	lines := out.ReadLines(len(OUTPUT))

	for i, line := range lines {
		assert.Equal(t, OUTPUT[i], line)
	}
}

func TestStdoutPipePrintMany(t *testing.T) {
	//-- arrange
	const SEED = 42
	r := rand.New(SEED)

	OUTPUT := []string{r.ASCII(15), r.ASCII(15), r.ASCII(15)}

	out := pipe.OpenStdout(1)
	defer out.Close()

	for _, output := range OUTPUT {
		//-- act
		fmt.Println(output)

		//-- assert
		assert.Equal(t, output, out.ReadLine())
	}
}
