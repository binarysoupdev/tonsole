package testio

import (
	"bufio"
	"io"
	"os"
)

// StdoutPipe is a test helper for capturing stdout.
// The primary use case is to test program output typically printed to the console.
type StdoutPipe struct {
	stdout *os.File
	output io.ReadCloser
	buffer chan string
}

// Create a new StdoutPipe. After calling, stdout will write output to the pipe until Close is called.
func OpenStdoutPipe() StdoutPipe {
	p := StdoutPipe{
		stdout: os.Stdout,
		buffer: make(chan string),
	}
	p.output, os.Stdout, _ = os.Pipe()

	go p.run()
	return p
}

// Close the pipe and restore stdout. The pipe can no longer be written to or read from.
func (p StdoutPipe) Close() {
	p.output.Close()
	os.Stdout.Close()

	os.Stdout = p.stdout
}

// Read the next line in the pipe. If there are no more lines, the execution stalls until output is available.
func (p StdoutPipe) NextLine() string {
	return <-p.buffer
}

func (p StdoutPipe) run() {
	scanner := bufio.NewScanner(p.output)

	for scanner.Scan() {
		p.buffer <- scanner.Text()
	}
}
