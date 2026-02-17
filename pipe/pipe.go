package pipe

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Pair struct {
	Prompt string
	Value  any
}

type IOPipe struct {
	stdin  *os.File
	stdout *os.File

	input   io.WriteCloser
	output  io.ReadCloser
	scanner *bufio.Scanner

	buffer chan Pair
	done   chan struct{}
}

func OpenStdio(bufSize int) IOPipe {
	p := IOPipe{
		stdin:  os.Stdin,
		stdout: os.Stdout,
		buffer: make(chan Pair, bufSize),
		done:   make(chan struct{}),
	}

	os.Stdin, p.input, _ = os.Pipe()
	p.output, os.Stdout, _ = os.Pipe()
	p.scanner = bufio.NewScanner(p.output)

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

	p.done <- struct{}{}
}

func (p IOPipe) Submit(pairs ...Pair) {
	for _, pair := range pairs {
		p.buffer <- pair
	}
}

func (p IOPipe) run() {
	for {
		select {
		case <-p.done:
			return
		case pair := <-p.buffer:
			p.queueInput(pair)
		}
	}
}

func (p IOPipe) queueInput(pair Pair) {
	for p.scanner.Scan() {
		if strings.Contains(p.scanner.Text(), pair.Prompt) {
			fmt.Fprintln(p.input, pair.Value)
			return
		}
	}
}
