package tinsel_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/binarysoupdev/tinsel/rand"
	"github.com/binarysoupdev/tinsel/tinsel"
	"github.com/stretchr/testify/assert"
)

func TestStdoutPipePrintOnce(t *testing.T) {
	//-- arrange
	r := rand.New(SEED)
	OUTPUT := []string{r.ASCII(10), r.ASCII(10), r.ASCII(10)}

	out := tinsel.OpenStdoutPipe()
	defer out.Close()

	//-- act
	fmt.Println(strings.Join(OUTPUT, "\n"))

	for i := range OUTPUT {
		//-- assert
		assert.Equal(t, OUTPUT[i], out.ReadLine())
	}
}

func TestStdoutPipePrintMany(t *testing.T) {
	//-- arrange
	r := rand.New(SEED)
	OUTPUT := []string{r.ASCII(10), r.ASCII(10), r.ASCII(10)}

	out := tinsel.OpenStdoutPipe()
	defer out.Close()

	for _, output := range OUTPUT {
		//-- act
		fmt.Println(output)

		//-- assert
		assert.Equal(t, output, out.ReadLine())
	}
}
