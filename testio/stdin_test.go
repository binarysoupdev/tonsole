package testio_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/binarysoupdev/tinsel/rand"
	"github.com/binarysoupdev/tinsel/testio"
	"github.com/stretchr/testify/assert"
)

func TestStdinPipe(t *testing.T) {
	//-- arrange
	r := rand.New(64)
	INPUT := []any{r.ASCII(10), r.ASCII(10), r.ASCII(10)}

	in := testio.OpenStdinPipe(len(INPUT))
	defer in.Close()

	//-- act
	in.Submit(INPUT...)

	for i := range INPUT {
		testio.Notify()
		res, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		//-- assert
		assert.Equal(t, INPUT[i], res[:len(res)-1])
	}
}
