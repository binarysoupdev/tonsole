package pipe

import (
	"io"
	"os"
)

type StdioPipe struct {
	stdin  *os.File
	stdout *os.File

	input  io.WriteCloser
	output io.ReadCloser

	inBuffer    chan InputPair
	outBuffer   chan string
	inputClosed chan struct{}
	close       chan struct{}

	echo bool
}

func OpenStdio(inputBuf, outputBuf int, echo bool) StdioPipe {
	p := StdioPipe{
		stdin:       os.Stdin,
		stdout:      os.Stdout,
		input:       nil,
		inBuffer:    nil,
		outBuffer:   nil,
		inputClosed: make(chan struct{}, 1),
		close:       make(chan struct{}, 1),
		echo:        echo,
	}

	if inputBuf > 0 {
		os.Stdin, p.input, _ = os.Pipe()
		p.inBuffer = make(chan InputPair, inputBuf)
	}

	p.output, os.Stdout, _ = os.Pipe()
	if outputBuf > 0 {
		p.outBuffer = make(chan string, outputBuf)
	}

	go p.run()
	return p
}

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
