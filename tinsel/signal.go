package tinsel

import "fmt"

var signal = make(chan struct{})

// Signal the stdin pipe to queue the next input.
// Should be called right before reading from stdin.
//
// Optionally print a newline to stdout to mimic pressing enter in a terminal.
//
// Note: if the pipe is not open, this function does nothing.
func QueueInput(newline bool) {
	if input == nil {
		return
	}

	if newline {
		fmt.Println()
	}
	signal <- struct{}{}
}
