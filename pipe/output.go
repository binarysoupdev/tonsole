package pipe

import "bufio"

func OpenStdout(bufSize int) IOPipe {
	return OpenStdio(0, bufSize, false)
}

func (p IOPipe) ReadLine() string {
	return <-p.outBuffer
}

func (p IOPipe) ReadLines(count int) []string {
	lines := make([]string, count)

	for i := range count {
		lines[i] = p.ReadLine()
	}
	return lines
}

func (p IOPipe) outputLoop() {
	if p.outBuffer == nil {
		return
	}
	scanner := bufio.NewScanner(p.output)

	for scanner.Scan() {
		p.outBuffer <- scanner.Text()
	}
}
