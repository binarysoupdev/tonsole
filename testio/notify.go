package testio

var signal = make(chan struct{})

// Notify the goroutine to queue the next input.
// Should be called right before reading from stdin.
//
// Note: if the StdinPipe is not open, this function does nothing.
func Notify() {
	if input != nil {
		signal <- struct{}{}
	}
}
