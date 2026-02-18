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

const SEED = 42

func TestStdioPipeInputAndOutput(t *testing.T) {
	//-- arrange
	r := rand.New(SEED)
	INPUT := pipe.InputPair{"prompt", r.ASCII(10)}
	PRE_INPUT := r.ASCII(15)
	POST_INPUT := r.ASCII(15)

	io := pipe.OpenStdio(1, 3)
	defer io.Close()

	io.SubmitFinal(INPUT)

	//-- act
	fmt.Println(PRE_INPUT)

	fmt.Println(INPUT.Prompt + ": ")
	res, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println(POST_INPUT)

	//-- assert
	assert.Equal(t, INPUT.Value, res[:len(res)-1])

	assert.Equal(t, PRE_INPUT, io.ReadLine())
	assert.Contains(t, io.ReadLine(), INPUT.Prompt)
	assert.Equal(t, POST_INPUT, io.ReadLine())
}
