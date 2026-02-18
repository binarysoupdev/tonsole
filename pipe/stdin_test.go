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

func TestStdinPipeSubmitOnce(t *testing.T) {
	//-- arrange
	r := rand.New(SEED)
	INPUT := []pipe.InputPair{{"prompt1: ", r.ASCII(10)}, {"prompt2: ", r.ASCII(10)}, {"prompt3: ", r.ASCII(10)}}

	in := pipe.OpenStdin(len(INPUT))
	defer in.Close()

	//-- act
	in.Submit(INPUT...)

	for _, input := range INPUT {
		fmt.Print(input.Prompt)
		res := readStdin()

		//-- assert
		assert.Equal(t, input.Value, res)
	}
}

func TestStdinPipeSubmitMany(t *testing.T) {
	//-- arrange
	r := rand.New(SEED)
	INPUT := []pipe.InputPair{{"prompt1: ", r.ASCII(10)}, {"prompt2: ", r.ASCII(10)}, {"prompt3: ", r.ASCII(10)}}

	in := pipe.OpenStdin(1)
	defer in.Close()

	for _, input := range INPUT {
		//-- act
		in.Submit(input)

		fmt.Print(input.Prompt)
		res := readStdin()

		//-- assert
		assert.Equal(t, input.Value, res)
	}
}

func TestStdinPipeReadPrompt(t *testing.T) {
	//-- arrange
	r := rand.New(SEED)
	INPUT := pipe.InputPair{"prompt", r.ASCII(10)}

	in := pipe.OpenStdin(1)
	defer in.Close()

	in.Submit(INPUT, INPUT)

	//-- act
	fmt.Print("---prompt")
	res1 := readStdin()

	fmt.Print("pro-pro-prompt")
	res2 := readStdin()

	//-- assert
	assert.Equal(t, INPUT.Value, res1)
	assert.Equal(t, INPUT.Value, res2)
}
