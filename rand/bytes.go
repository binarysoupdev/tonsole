package rand

// Generate a random sequence of bytes.
func (r Rand) Bytes(len int) []byte {
	bytes := make([]byte, len)
	r.Read(bytes)

	return bytes
}

// Generate a random ASCII string of printable characters.
func (r Rand) ASCII(len int) string {
	bytes := make([]byte, len)
	for i := range bytes {
		// printable ASCII range
		bytes[i] = byte(r.IntRange(32, 126))
	}
	return string(bytes)
}
