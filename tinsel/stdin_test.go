package tinsel_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/binarysoupdev/tinsel/rand"
	"github.com/binarysoupdev/tinsel/tinsel"
	"github.com/stretchr/testify/assert"
)

func TestStdinPipeSubmitOnce(t *testing.T) {
	//-- arrange
	r := rand.New(64)
	INPUT := []any{r.ASCII(10), r.ASCII(10), r.ASCII(10)}

	in := tinsel.OpenStdinPipe(len(INPUT))
	defer in.Close()

	//-- act
	in.Submit(INPUT...)

	for i := range INPUT {
		tinsel.QueueInput()
		res, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		//-- assert
		assert.Equal(t, INPUT[i], res[:len(res)-1])
	}
}

func TestStdinPipeSubmitMany(t *testing.T) {
	//-- arrange
	r := rand.New(64)
	INPUT := []any{r.ASCII(10), r.ASCII(10), r.ASCII(10)}

	in := tinsel.OpenStdinPipe(len(INPUT))
	defer in.Close()

	for _, input := range INPUT {
		//-- act
		in.Submit(input)

		tinsel.QueueInput()
		res, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		//-- assert
		assert.Equal(t, input, res[:len(res)-1])
	}
}
