package tinsel

import (
	"bufio"
	"io"
	"os"
)

// StdoutPipe is a test helper for capturing stdout.
// The primary use case is to test program output typically printed to the console.
type StdoutPipe struct {
	stdout  *os.File
	output  io.ReadCloser
	scanner *bufio.Scanner
}

// Create a new StdoutPipe. After calling, stdout will write output to the pipe until Close is called.
func OpenStdoutPipe() StdoutPipe {
	p := StdoutPipe{
		stdout: os.Stdout,
	}
	p.output, os.Stdout, _ = os.Pipe()

	p.scanner = bufio.NewScanner(p.output)
	return p
}

// Close the pipe and restore stdout. The pipe can no longer be written to or read from.
func (p StdoutPipe) Close() {
	p.output.Close()
	os.Stdout.Close()

	os.Stdout = p.stdout
}

// Read from the pipe until a newline is encountered.
// If there are no newlines in the pipe, the program stalls until one is written (see EndLine).
func (p StdoutPipe) ReadLine() string {
	if p.scanner.Scan() {
		return p.scanner.Text()
	} else {
		panic("stdout pipe is closed")
	}
}
