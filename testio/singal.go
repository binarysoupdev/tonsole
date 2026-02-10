package testio

import (
	"fmt"
	"io"
)

var signal = make(chan struct{})
var input io.WriteCloser

// Notify the go routine to queue the next input.
// Should be called right before reading from stdin.
//
// If the StdinPipe is not open, this function does nothing.
func Notify() {
	if input != nil {
		signal <- struct{}{}
	}
}

func run(buffer chan any) {
	for {
		<-signal
		fmt.Fprintln(input, <-buffer)
	}
}
