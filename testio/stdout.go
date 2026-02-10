package testio

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// StdoutPipe is a test helper for capturing stdout.
// The primary use case is to test program output typically printed to the console.
type StdoutPipe struct {
	stdout  *os.File
	in      *os.File
	out     *os.File
	scanner *bufio.Scanner
}

// Create a new StdoutPipe. After calling, stdout will write output to the pipe until Close is called.
func OpenStdoutPipe() StdoutPipe {
	p := StdoutPipe{
		stdout: os.Stdout,
	}
	p.out, p.in, _ = os.Pipe()

	os.Stdout = p.in
	p.scanner = bufio.NewScanner(p.out)
	return p
}

// Close the input. The pipe can still be read from, but will no longer be written to.
//
// Should be called before reading to avoid unexpected program stalls.
func (p StdoutPipe) CloseInput() {
	p.in.Close()
}

// Close the pipe and restore stdout. The pipe can no longer be written to or read from.
func (p StdoutPipe) Close() {
	os.Stdout = p.stdout
	p.in.Close()
	p.out.Close()
}

// Read the next line in the pipe. If there are no more lines, the test fails.
func (p StdoutPipe) NextLine(t *testing.T) string {
	if p.scanner.Scan() {
		return p.scanner.Text()
	}

	assert.FailNow(t, "reached end of stdout pipe")
	return ""
}

// Assert the pipe has no more lines.
func (p StdoutPipe) TestEOF(t *testing.T) {
	assert.False(t, p.scanner.Scan(), "stdout pipe has more lines")
}
