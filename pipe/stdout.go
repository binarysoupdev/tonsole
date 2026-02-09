package pipe

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

// Create a new StdoutPipe. Note: the old stdout is cached at create time, not when opened.
func Stdout() StdoutPipe {
	return StdoutPipe{
		stdout: os.Stdout,
	}
}

// Start capturing output from stdout.
//
// Returns a copy of the pipe to enable function chaining.
func (p *StdoutPipe) Open() StdoutPipe {
	p.out, p.in, _ = os.Pipe()

	os.Stdout = p.in
	p.scanner = bufio.NewScanner(p.out)

	return *p
}

// Close the input and restore stdout.
// The pipe can still be read from, but will no longer be written to.
//
// Note: closing the input before reading may be required to avoid some errors.
func (p StdoutPipe) CloseInput() {
	os.Stdout = p.stdout
	p.in.Close()
}

// Close the pipe and restore the cached stdout.
// The pipe can no longer be written to or read from.
func (p StdoutPipe) Close() {
	p.CloseInput()
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
