/*
The pipe package provides a method to input values to stdin and read output from stdout in automated test cases.
The following example demonstrates a basic usage case for the pipe.

	// create new prompt and input pair
	INPUT := pipe.InputPair{"prompt: ", "input"}

	// open the pipe with buffer sizes of 1 and echo enabled
	io := pipe.OpenStdio(1, 1, true)
	defer io.Close()

	// queue the input and mark input as complete
	io.QueueFinal(INPUT)

	// write the prompt to stdout then read input from stdin
	fmt.Print(INPUT.Prompt)
	res, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// (using testify) assert the result matches the input with newline stripped
	assert.Equal(t, INPUT.Value, res[:len(res)-1])

	// assert the next line of the output contains the prompt string and its value
	assert.Equal(t, fmt.Sprintf("%s%v", INPUT.Prompt, INPUT.Value), io.ReadLine())
*/
package pipe
