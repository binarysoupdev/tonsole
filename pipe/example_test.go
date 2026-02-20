package pipe_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/binarysoupdev/tinsel/pipe"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/assert"
)

func TestStdioPipeBasicExample(t *testing.T) {
	// create rand object with constant seed
	const SEED = 42
	r := rand.New(SEED)

	// create new random input string
	INPUT := r.ASCII(15)

	// open the pipe with buffer sizes of 1 and echo enabled
	io := pipe.OpenStdio(1, 1, true)
	defer io.Close()

	// queue the input with the expected prompt and mark the input sequence as complete
	io.Queue("prompt: ", INPUT)
	io.EndQueue()

	// read the input from stdin
	fmt.Print("prompt: ")
	res, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// assert the result matches the input (with newline stripped)
	assert.Equal(t, INPUT, res[:len(res)-1])

	// assert the next line of the output contains the prompt string and its input
	assert.Equal(t, fmt.Sprintf("prompt: %s", INPUT), io.ReadLine())
}
