package testio_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/binarysoupdev/tinsel/rand"
	"github.com/binarysoupdev/tinsel/testio"
	"github.com/stretchr/testify/assert"
)

func TestStdoutPipe(t *testing.T) {
	//-- arrange
	r := rand.New(64)
	OUTPUT := []string{r.ASCII(10), r.ASCII(10), r.ASCII(10)}

	out := testio.OpenStdoutPipe()
	defer out.Close()

	//-- act
	fmt.Println(strings.Join(OUTPUT, "\n"))

	for i := range OUTPUT {
		//-- assert
		assert.Equal(t, OUTPUT[i], out.NextLine())
	}
}
