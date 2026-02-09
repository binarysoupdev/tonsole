package rand

// Generate a random int between [min, max] inclusive.
func (r Rand) IntRange(min, max int) int {
	return r.Intn(max-min+1) + min
}
