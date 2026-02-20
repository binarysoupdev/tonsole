/*
The pipe package provides the StdioPipe for reading/writing from stdin/stdout in automated testing.

The create a new pipe, use the Open method with the desired buffer sizes.

	// open pipe then defer close
	io := pipe.OpenStdio(1, 1, true)
	defer io.Close()

If only input or output is desired, use the appropriate open variant.

	// stdin only
	in := pipe.OpenStdin(1)
	defer in.Close()

	// stdout only
	out := pipe.OpenStdout(1)
	defer out.Close()

To queue input for stdin, supply the input with the expected prompt using the Queue method.

	// queue a [prompt, input] pair
	in.Queue("prompt: ", INPUT)

If using both input and output (ie. OpenStdio), then the queue sequence must be marked as complete.

	// queue the pair then mark finished
	io.Queue("prompt: ", INPUT)
	io.EndQueue()

To read output from stdout (newline separated), use the ReadLine or ReadLines variant.

	// read single line
	line := out.ReadLine()

	// read multiple lines (returns slice)
	lines := out.ReadLines(2)

To read a line, but ignore it's value, use SkipLines.

	// skip the following lines
	out.SkipLines(2)
*/
package pipe
