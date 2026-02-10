package testio

import (
	"fmt"
	"io"
	"os"
)

var input io.WriteCloser

// StdinPipe is a test helper for inputting stdin.
// The primary use case is to submit program input typically collected from the console.
type StdinPipe struct {
	stdin  *os.File
	buffer chan any
	done   chan struct{}
}

// Create a new StdinPipe. After calling, stdin will read input from the pipe until Close is called.
//
// 'bufSize' defines the size of the input buffer. Ensure the size is large enough to prevent stalling.
func OpenStdinPipe(bufSize int) StdinPipe {
	p := StdinPipe{
		stdin:  os.Stdin,
		buffer: make(chan any, bufSize),
		done:   make(chan struct{}),
	}
	os.Stdin, input, _ = os.Pipe()

	go p.run()
	return p
}

// Close the pipe and restore stdin. The pipe can no longer be read from or written to.
func (p StdinPipe) Close() {
	os.Stdin.Close()
	input.Close()

	os.Stdin = p.stdin
	input = nil

	p.done <- struct{}{}
}

// Submit new input to the buffer. If the buffer is full, the execution will stall until room is made.
func (p StdinPipe) Submit(input ...any) {
	for _, val := range input {
		p.buffer <- val
	}
}

func (p StdinPipe) run() {
	for {
		select {
		case <-p.done:
			return
		case <-signal:
			fmt.Fprintln(input, <-p.buffer)
		}
	}
}
