package rand

import "math/rand"

// Test helper for generating random test values.
type Rand struct {
	*rand.Rand
}

// Create a new rand object from a seed. A static seed should be used to ensure deterministic tests.
func New(seed int64) Rand {
	return Rand{
		Rand: rand.New(rand.NewSource(seed)),
	}
}
