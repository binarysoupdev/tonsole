package pipe

import (
	"fmt"
	"strings"
)

type InputPair struct {
	Prompt string
	Value  any
}

func OpenStdin(bufSize int) IOPipe {
	return OpenStdio(bufSize, 0, false)
}

func (p IOPipe) Submit(pairs ...InputPair) {
	for _, pair := range pairs {
		p.inBuffer <- pair
	}
}

func (p IOPipe) SubmitFinal(pairs ...InputPair) {
	p.Submit(pairs...)
	p.inputClosed <- struct{}{}
}

func (p IOPipe) inputLoop() {
	if p.inBuffer == nil {
		return
	}
	closed := false

	for {
		select {
		case <-p.cancel:
			return
		case pair := <-p.inBuffer:
			p.inputOnPrompt(pair.Prompt, pair.Value)
		}

		if len(p.inputClosed) > 0 {
			<-p.inputClosed
			closed = true
		}

		if closed && len(p.inBuffer) == 0 {
			return
		}
	}
}

func (p IOPipe) inputOnPrompt(prompt string, val any) {
	output := ""
	p.waitPrompt([]byte(prompt), &output)

	if p.echo {
		p.captureOutput(&output, fmt.Sprintf("%v\n", val))
	} else {
		p.captureOutput(&output, "\n")
	}

	fmt.Fprintln(p.input, val)
}

func (p IOPipe) waitPrompt(prompt []byte, output *string) {
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

func (p IOPipe) captureOutput(output *string, str string) {
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
