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

func TestStdioPipeSubmitOnce(t *testing.T) {
	//-- arrange
	r := rand.New(SEED)
	INPUT := []pipe.Pair{{"prompt1", r.ASCII(10)}, {"prompt2", r.ASCII(10)}, {"prompt3", r.ASCII(10)}}

	io := pipe.OpenStdio(len(INPUT))
	defer io.Close()

	//-- act
	io.Submit(INPUT...)

	for _, input := range INPUT {
		fmt.Println(input.Prompt + ": ")
		res, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		//-- assert
		assert.Equal(t, input.Value, res[:len(res)-1])
	}
}

func TestStdioPipeSubmitMany(t *testing.T) {
	//-- arrange
	r := rand.New(SEED)
	INPUT := []pipe.Pair{{"prompt1", r.ASCII(10)}, {"prompt2", r.ASCII(10)}, {"prompt3", r.ASCII(10)}}

	io := pipe.OpenStdio(1)
	defer io.Close()

	for _, input := range INPUT {
		//-- act
		io.Submit(input)

		fmt.Println(input.Prompt + ": ")
		res, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		//-- assert
		assert.Equal(t, input.Value, res[:len(res)-1])
	}
}
