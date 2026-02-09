package pipe

import (
	"fmt"
	"os"
)

// StdinPipe is a test helper for inputting stdin.
// The primary use case is to submit program input typically collected from the console.
type StdinPipe struct {
	stdin *os.File
	out   *os.File
	in    *os.File
}

// Create a new StdinPipe. Note: the old stdin is cached at create time, not when opened.
func Stdin() StdinPipe {
	return StdinPipe{
		stdin: os.Stdin,
	}
}

// Start capturing input from stdin. Returns a copy of the pipe to enable function chaining.
func (p *StdinPipe) Open(input []any) StdinPipe {
	p.out, p.in, _ = os.Pipe()

	os.Stdin = p.out
	return *p
}

// Close the output and restore stdin.
// The pipe can still be written to, but will no longer be read from.
//
// Has little practical use if called explicitly, but is internally called by Close.
func (p StdinPipe) CloseOutput() {
	os.Stdin = p.stdin
	p.out.Close()
}

// Close the pipe and restore stdin.
// The pipe can no longer be read from or written to.
func (p StdinPipe) Close() {
	p.CloseOutput()
	p.in.Close()
}

// Write the input as newline separated tokens to the pipe.
func (p StdinPipe) WriteLines(input []any) {
	for _, line := range input {
		fmt.Fprintln(p.in, line)
	}
}
