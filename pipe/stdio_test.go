package pipe_test

import (
	"fmt"
	"testing"

	"github.com/binarysoupdev/tinsel/pipe"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/assert"
)

func TestStdioPipeInputAndOutput(t *testing.T) {
	//-- arrange
	const SEED = 42
	r := rand.New(SEED)

	INPUT := pipe.InputPair{"prompt: ", r.ASCII(10)}
	PRE_INPUT := r.ASCII(15)
	POST_INPUT := r.ASCII(15)

	io := pipe.OpenStdio(1, 3, false)
	defer io.Close()

	io.QueueFinal(INPUT)

	//-- act
	fmt.Println(PRE_INPUT)

	fmt.Print(INPUT.Prompt)
	res := readStdin()

	fmt.Println(POST_INPUT)

	//-- assert
	assert.Equal(t, INPUT.Value, res)

	assert.Equal(t, PRE_INPUT, io.ReadLine())
	assert.Contains(t, io.ReadLine(), INPUT.Prompt)
	assert.Equal(t, POST_INPUT, io.ReadLine())
}

func TestStdioPipeWithEcho(t *testing.T) {
	//-- arrange
	const SEED = 42
	r := rand.New(SEED)

	INPUT := pipe.InputPair{"prompt: ", r.ASCII(10)}

	io := pipe.OpenStdio(1, 1, true)
	defer io.Close()

	io.QueueFinal(INPUT)

	//-- act
	fmt.Print(INPUT.Prompt)
	res := readStdin()

	//-- assert
	assert.Equal(t, INPUT.Value, res)
	assert.Equal(t, fmt.Sprintf("%s%v", INPUT.Prompt, INPUT.Value), io.ReadLine())
}
