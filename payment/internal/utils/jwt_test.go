package utils

import (
	"testing"
)

func BenchmarkTokenPair(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = TokenPair(randomString(b))
	}
}
