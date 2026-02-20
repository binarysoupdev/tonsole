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

	PROMPT := r.ASCII(10)
	INPUT := r.ASCII(15)

	PRE_INPUT := r.ASCII(30)
	POST_INPUT := r.ASCII(30)

	io := pipe.OpenStdio(1, 3, false)
	defer io.Close()

	io.Queue(PROMPT, INPUT)
	io.EndQueue()

	//-- act
	fmt.Println(PRE_INPUT)

	fmt.Print(PROMPT)
	res := readStdin()

	fmt.Println(POST_INPUT)

	//-- assert
	assert.Equal(t, INPUT, res)

	assert.Equal(t, PRE_INPUT, io.ReadLine())
	assert.Contains(t, io.ReadLine(), PROMPT)
	assert.Equal(t, POST_INPUT, io.ReadLine())
}

func TestStdioPipeWithEcho(t *testing.T) {
	//-- arrange
	const SEED = 42
	r := rand.New(SEED)

	PROMPT := r.ASCII(10)
	INPUT := r.ASCII(15)

	io := pipe.OpenStdio(1, 1, true)
	defer io.Close()

	io.Queue(PROMPT, INPUT)
	io.EndQueue()

	//-- act
	fmt.Print(PROMPT)
	res := readStdin()

	//-- assert
	assert.Equal(t, INPUT, res)
	assert.Equal(t, fmt.Sprintf("%s%v", PROMPT, INPUT), io.ReadLine())
}
