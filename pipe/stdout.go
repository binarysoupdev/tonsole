package pipe

import "bufio"

// Open just the stdout pipe. Equivalent to a StdioPipe with no input buffer.
func OpenStdout(bufSize int) StdioPipe {
	return OpenStdio(0, bufSize, false)
}

// Read the next line from the output buffer.
// Blocks until a line is available.
func (p StdioPipe) ReadLine() string {
	return <-p.outBuffer
}

// Read multiple lines at once from output buffer.
// Blocks until enough lines are available.
func (p StdioPipe) ReadLines(count int) []string {
	lines := make([]string, count)

	for i := range count {
		lines[i] = p.ReadLine()
	}
	return lines
}

func (p StdioPipe) outputLoop() {
	if p.outBuffer == nil {
		return
	}
	scanner := bufio.NewScanner(p.output)

	for scanner.Scan() {
		p.outBuffer <- scanner.Text()
	}
}
