package pipe

import (
	"fmt"
	"strings"
)

type inputPair struct {
	Prompt string
	Value  any
}

// Open just the stdin pipe. Equivalent to a StdioPipe with no output buffer.
func OpenStdin(bufSize int) StdioPipe {
	return OpenStdio(bufSize, 0, false)
}

// Queue the next [prompt, input] pair to the buffer.
// Stalls if the input buffer is full (ensure a larger enough buffer size).
//
// The pipe will read stdout until the prompt has been read exactly,
// then the associated input will be written to stdin.
func (p StdioPipe) Queue(prompt string, input any) {
	p.inBuffer <- inputPair{prompt, input}
}

// Notify the pipe that no more input is expected.
// Must be called to continue reading output after the final input.
func (p StdioPipe) EndQueue() {
	p.inputClosed <- struct{}{}
}

func (p StdioPipe) inputLoop() {
	if p.inBuffer == nil {
		return
	}

	for len(p.inputClosed) == 0 || len(p.inBuffer) > 0 {
		select {
		case <-p.close:
			return
		case pair := <-p.inBuffer:
			p.inputOnPrompt(pair.Prompt, pair.Value)
		}
	}
}

func (p StdioPipe) inputOnPrompt(prompt string, val any) {
	output := ""
	p.waitPrompt([]byte(prompt), &output)

	if p.echo {
		p.captureOutput(&output, fmt.Sprintf("%v\n", val))
	} else {
		p.captureOutput(&output, "\n")
	}

	fmt.Fprintln(p.input, val)
}

func (p StdioPipe) waitPrompt(prompt []byte, output *string) {
	buffer := make([]byte, len(prompt))
	slice := buffer[:]
	index := 0

	for {
		p.output.Read(slice)
		p.captureOutput(output, string(slice))

		for _, char := range slice {
			if char == prompt[index] {
				index++
			} else {
				index = 0
			}
		}

		if index == len(prompt) {
			return
		}
		slice = buffer[:len(prompt)-index]
	}
}

func (p StdioPipe) captureOutput(output *string, str string) {
	if p.outBuffer == nil {
		return
	}

	*output += str
	tokens := strings.Split(*output, "\n")

	for i := range len(tokens) - 1 {
		p.outBuffer <- tokens[i]
	}
	*output = tokens[len(tokens)-1]
}
