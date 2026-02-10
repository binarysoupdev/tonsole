package testio_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/binarysoupdev/tinsel/rand"
	"github.com/binarysoupdev/tinsel/testio"
	"github.com/stretchr/testify/assert"
)

func TestStdinPipeSubmitOnce(t *testing.T) {
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

func TestStdinPipeSubmitMany(t *testing.T) {
	//-- arrange
	r := rand.New(64)
	INPUT := []any{r.ASCII(10), r.ASCII(10), r.ASCII(10)}

	in := testio.OpenStdinPipe(len(INPUT))
	defer in.Close()

	for _, input := range INPUT {
		//-- act
		in.Submit(input)

		testio.Notify()
		res, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		//-- assert
		assert.Equal(t, input, res[:len(res)-1])
	}
}
