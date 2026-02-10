package testio

import (
	"os"
)

// StdinPipe is a test helper for inputting stdin.
// The primary use case is to submit program input typically collected from the console.
type StdinPipe struct {
	stdin *os.File
	out   *os.File
	in    *os.File
}

// Create a new StdinPipe. After calling, stdin will read input from the pipe until Restore is called.
//
// To prevent program stalling, all input should be submitted to the pipe before it is needed.
func OpenStdinPipe() StdinPipe {
	p := StdinPipe{
		stdin: os.Stdin,
	}
	p.out, p.in, _ = os.Pipe()

	os.Stdin = p.out
	return p
}

// Close the pipe and restore stdin. The pipe can no longer be read from or written to.
func (p StdinPipe) Restore() {
	os.Stdin = p.stdin
	p.out.Close()
	p.in.Close()
}

// Submit input to the pipe. Should only be called once per test.
func (p StdinPipe) Submit(input ...any) {
	go queueInput(p.in, input)
}
