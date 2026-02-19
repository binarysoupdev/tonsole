package pipe

import (
	"io"
	"os"
)

// StdioPipe is a testing utility for inputting stdin and reading stdout in automated test cases.
type StdioPipe struct {
	stdin  *os.File
	stdout *os.File

	input  io.WriteCloser
	output io.ReadCloser

	inBuffer    chan inputPair
	outBuffer   chan string
	inputClosed chan struct{}
	close       chan struct{}

	echo bool
}

// Open the Stdio pipe with the requested buffer sizes.
// After calling, stdin/stdout will read/write to the pipe until Close is called.
// Do to the redirecting of stdin and stdout, usage of the StdioPipe is NOT thread safe.
//
// Note the process will stall if either buffer is too small to store the required values, so
// ensure a large enough buffer size to accommodate all expected data.
// If either buffer size is zero, then that buffer is effectively disabled (see OpenStdin/OpenStdout).
//
// Echoing involves copying input to the output buffer to mimic terminal echoing.
// Does nothing if either buffer is disabled.
func OpenStdio(inBufSize, outBufSize int, enableEcho bool) StdioPipe {
	p := StdioPipe{
		stdin:       os.Stdin,
		stdout:      os.Stdout,
		input:       nil,
		inBuffer:    nil,
		outBuffer:   nil,
		inputClosed: make(chan struct{}, 1),
		close:       make(chan struct{}, 1),
		echo:        enableEcho,
	}

	if inBufSize > 0 {
		os.Stdin, p.input, _ = os.Pipe()
		p.inBuffer = make(chan inputPair, inBufSize)
	}

	p.output, os.Stdout, _ = os.Pipe()
	if outBufSize > 0 {
		p.outBuffer = make(chan string, outBufSize)
	}

	go p.run()
	return p
}

// Close the pipe and restore stdin/stdout.
func (p StdioPipe) Close() {
	if p.input != nil {
		os.Stdin.Close()
		p.input.Close()
	}

	os.Stdout.Close()
	p.output.Close()

	os.Stdin = p.stdin
	os.Stdout = p.stdout

	p.close <- struct{}{}
}

func (p StdioPipe) run() {
	p.inputLoop()
	p.outputLoop()
}
