package pipe

import (
	"io"
	"os"
)

type IOPipe struct {
	stdin  *os.File
	stdout *os.File

	input  io.WriteCloser
	output io.ReadCloser

	inBuffer    chan InputPair
	outBuffer   chan string
	inputClosed chan struct{}
	cancel      chan struct{}

	echo bool
}

func OpenStdio(inputBuf, outputBuf int, echo bool) IOPipe {
	p := IOPipe{
		stdin:       os.Stdin,
		stdout:      os.Stdout,
		inBuffer:    nil,
		outBuffer:   nil,
		inputClosed: make(chan struct{}, 1),
		cancel:      make(chan struct{}, 1),
		echo:        echo,
	}

	os.Stdin, p.input, _ = os.Pipe()
	p.output, os.Stdout, _ = os.Pipe()

	if inputBuf > 0 {
		p.inBuffer = make(chan InputPair, inputBuf)
	}

	if outputBuf > 0 {
		p.outBuffer = make(chan string, outputBuf)
	}

	go p.run()
	return p
}

func (p IOPipe) Close() {
	os.Stdin.Close()
	p.input.Close()

	os.Stdout.Close()
	p.output.Close()

	os.Stdin = p.stdin
	os.Stdout = p.stdout

	p.cancel <- struct{}{}
}

func (p IOPipe) run() {
	p.inputLoop()
	p.outputLoop()
}
