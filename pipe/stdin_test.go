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

func readStdin() string {
	res, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return res[:len(res)-1]
}

func TestStdinPipeQueueAll(t *testing.T) {
	//-- arrange
	const SEED = 42
	r := rand.New(SEED)

	PROMPT := r.ASCII(10)
	INPUTS := []string{r.ASCII(15), r.ASCII(15), r.ASCII(15)}

	in := pipe.OpenStdin(len(INPUTS))
	defer in.Close()

	//-- act
	for _, input := range INPUTS {
		in.Queue(PROMPT, input)
	}

	for _, input := range INPUTS {
		fmt.Print(PROMPT)
		res := readStdin()

		//-- assert
		assert.Equal(t, input, res)
	}
}

func TestStdinPipeQueueOne(t *testing.T) {
	//-- arrange
	const SEED = 42
	r := rand.New(SEED)

	PROMPT := r.ASCII(10)
	INPUTS := []string{r.ASCII(15), r.ASCII(15), r.ASCII(15)}

	in := pipe.OpenStdin(1)
	defer in.Close()

	for _, input := range INPUTS {
		//-- act
		in.Queue(PROMPT, input)

		fmt.Print(PROMPT)
		res := readStdin()

		//-- assert
		assert.Equal(t, input, res)
	}
}

func TestStdinPipeReadPrompt(t *testing.T) {
	//-- arrange
	const SEED = 42
	r := rand.New(SEED)

	INPUT1 := r.ASCII(15)
	INPUT2 := r.ASCII(15)

	in := pipe.OpenStdin(1)
	defer in.Close()

	//-- act
	in.Queue("prompt", INPUT1)
	fmt.Print("---prompt")
	res1 := readStdin()

	in.Queue("prompt", INPUT2)
	fmt.Print("pro-pro-prompt")
	res2 := readStdin()

	//-- assert
	assert.Equal(t, INPUT1, res1)
	assert.Equal(t, INPUT2, res2)
}
