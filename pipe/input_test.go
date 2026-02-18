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

func TestStdinPipeSubmitOnce(t *testing.T) {
	//-- arrange
	r := rand.New(SEED)
	INPUT := []pipe.InputPair{{"prompt1", r.ASCII(10)}, {"prompt2", r.ASCII(10)}, {"prompt3", r.ASCII(10)}}

	in := pipe.OpenStdin(len(INPUT))
	defer in.Close()

	//-- act
	in.Submit(INPUT...)

	for _, input := range INPUT {
		fmt.Println(input.Prompt + ": ")
		res, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		//-- assert
		assert.Equal(t, input.Value, res[:len(res)-1])
	}
}

func TestStdinPipeSubmitMany(t *testing.T) {
	//-- arrange
	r := rand.New(SEED)
	INPUT := []pipe.InputPair{{"prompt1", r.ASCII(10)}, {"prompt2", r.ASCII(10)}, {"prompt3", r.ASCII(10)}}

	in := pipe.OpenStdin(1)
	defer in.Close()

	for _, input := range INPUT {
		//-- act
		in.Submit(input)

		fmt.Println(input.Prompt + ": ")
		res, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		//-- assert
		assert.Equal(t, input.Value, res[:len(res)-1])
	}
}
