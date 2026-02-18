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
	return OpenStdio(bufSize, 0)
}

func (p IOPipe) Submit(pairs ...InputPair) {
	for _, pair := range pairs {
		p.inBuffer <- pair
	}
}

func (p IOPipe) SubmitFinal(pairs ...InputPair) {
	p.Submit(pairs...)
	p.inputClosed = true
}

func (p IOPipe) inputLoop() {
	if p.inBuffer == nil {
		return
	}

	for len(p.inBuffer) > 0 || !p.inputClosed {
		select {
		case <-p.cancel:
			return
		case pair := <-p.inBuffer:
			p.queueInput(pair)
		}
	}
}

func (p IOPipe) queueInput(pair InputPair) {
	for p.scanner.Scan() {
		text := p.scanner.Text()

		if p.outBuffer != nil {
			p.outBuffer <- text
		}

		if strings.Contains(text, pair.Prompt) {
			fmt.Fprintln(p.input, pair.Value)
			return
		}
	}
}
