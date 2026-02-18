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
	for p.scanner.Scan() {
		text := p.scanner.Text()

		if p.outBuffer != nil {
			p.outBuffer <- text
		}

		if strings.Contains(text, prompt) {
			fmt.Fprintln(p.input, val)
			return
		}
	}
}
