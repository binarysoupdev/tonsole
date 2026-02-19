/*
The pipe package provides the ability to input values to stdin and read output from stdout in automated test cases.
The following example demonstrates a basic usage case for the pipe.

	// create new prompt and input pair
	PROMPT := "prompt: "
	INPUT := "input"

	// open the pipe with buffer sizes of 1 and echo enabled
	io := pipe.OpenStdio(1, 1, true)
	defer io.Close()

	// queue the input and mark the input sequence as complete
	io.Queue(INPUT)
	io.EndQueue()

	// write the prompt to stdout then read input from stdin
	fmt.Print(PROMPT)
	res, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// (using testify) assert the result matches the input with newline stripped
	assert.Equal(t, INPUT, res[:len(res)-1])

	// assert the next line of the output contains the prompt string and its value
	assert.Equal(t, fmt.Sprintf("%s%v", PROMPT, INPUT), io.ReadLine())
*/
package pipe
