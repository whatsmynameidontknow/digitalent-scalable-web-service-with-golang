package generator

import (
	"math/rand/v2"
	"testing"
)

func BenchmarkGenerateRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomString(1000, 15)
	}
}

func TestGenerateRandomString(t *testing.T) {
	for range 20 {
		numCh := 1000 + rand.IntN(1000)
		out := GenerateRandomString(numCh, 10)
		if len(out) != numCh {
			t.Errorf("expected generated string to have %d chars, got %d\n", numCh, len(out))
		}
	}
}
