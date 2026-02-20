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

// Read multiple lines at once from output buffer (see ReadLine).
func (p StdioPipe) ReadLines(count int) []string {
	lines := make([]string, count)

	for i := range count {
		lines[i] = p.ReadLine()
	}
	return lines
}

// Read then discard multiple lines from output buffer (see ReadLine).
func (p StdioPipe) SkipLines(count int) {
	for range count {
		_ = p.ReadLine()
	}
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
