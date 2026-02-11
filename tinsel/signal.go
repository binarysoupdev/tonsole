package tinsel

var signal = make(chan struct{})

// Signal the stdin pipe to queue the next input.
// Should be called right before reading from stdin.
//
// Note: if the pipe is not open, this function does nothing.
func QueueInput() {
	if input != nil {
		signal <- struct{}{}
	}
}
